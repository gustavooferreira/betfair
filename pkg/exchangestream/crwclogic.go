package exchangestream

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/gustavooferreira/betfair/pkg/globals"
	"github.com/gustavooferreira/betfair/pkg/utils/log"
)

// Timeout used by reader and writer
// timeout of 0.5 seconds
const timeoutDuration = 500 * time.Millisecond

func controller(esaclient *ESAClient, connMsgChan chan<- ConnectionMessage) {
	connPhaseDone := false

	respMsgChan := make(chan ResponseMessage, 1000)
	readerStopChan := make(chan bool)
	reqMsgChan := make(chan RequestMessage, 1000)
	writerStopInformChan := make(chan bool)
	connTrackerStopChan := make(chan bool)
	connTrackerStopInformChan := make(chan bool)

	// LookupTable (key: msgID)
	// TODO: implement expiration (if we don't get a response back, it never gets cleared!)
	lookupTable := make(map[uint32](chan ResponseMessage))

	// Spawn reader, writer and connTracker goroutines
	go esaclient.reader(respMsgChan, readerStopChan)
	go esaclient.writer(reqMsgChan, writerStopInformChan)
	go esaclient.connTracker(connTrackerStopChan, connTrackerStopInformChan)

	// If we get a signal on stopChan, trigger reader, writer and connTracker to end too
	// When getting the connection message, send the data and close the channel
	for {
		select {
		case <-esaclient.stopChan:
			close(connTrackerStopChan)
			close(readerStopChan)
			close(reqMsgChan)

			// Wait on all 3 informs and then return!
			// When getting all informs, send signal on stopInformChan Channel
			esaclient.stopInformChan <- true

			return
		case respMsg, ok := <-respMsgChan:
			if !ok {
				// This should never happen!
				// 	close(esaclient.stopInformChan)
				// 	return
			}
			// log.Log(globals.Logger, log.DEBUG, "betfair message received", log.Fields{"message": fmt.Sprintf("%+v", respMsg)})

			if respMsg.Op == "connection" {
				if !connPhaseDone {
					connMsgChan <- *respMsg.ConnectionMessage
					close(connMsgChan)
					connPhaseDone = true
				} else {
					log.Log(globals.Logger, log.ERROR, "got a ConnectionMessage while not being in connection phase", log.Fields{"connectionID": respMsg.ID})
				}
			} else if respMsg.Op == "mcm" {
				esaclient.MCMChan <- MarketChangeM{ID: respMsg.ID, MarketChangeMessage: *respMsg.MarketChangeMessage}
			} else if respMsg.Op == "ocm" {
				esaclient.OCMChan <- OrderChangeM{ID: respMsg.ID, OrderChangeMessage: *respMsg.OrderChangeMessage}
			} else if respMsg.Op == "status" {
				// Validate there is even an ID!
				if respMsg.ID == nil {
					// Error, no ID found!
					log.Log(globals.Logger, log.ERROR, "got status message without an ID", log.Fields{"message": fmt.Sprintf("%+v", respMsg)})
				} else if result, ok := lookupTable[*respMsg.ID]; !ok {
					// Error, no ID found!
					log.Log(globals.Logger, log.ERROR, "got status message with no matching ID", log.Fields{"message": fmt.Sprintf("%+v", respMsg)})
				} else {
					// Delete entry from the lookup table!
					delete(lookupTable, *respMsg.ID)

					result <- respMsg
				}
			} else {
				log.Log(globals.Logger, log.ERROR, "unknown message operation type", log.Fields{"message": fmt.Sprintf("%+v", respMsg)})
			}
		case workUnit, ok := <-esaclient.reqMsgChan:
			if !ok {
				// This should never happen!
				// Error
			}

			// Check if ID is set, if not, get one and set it on the struct
			if workUnit.req.ID == nil {
				temp := esaclient.getNewID()
				workUnit.req.ID = &temp
			} else if *workUnit.req.ID == 0 {
				*workUnit.req.ID = esaclient.getNewID()
			}

			// TODO: If message if of type MarketSub or OrderSub set heartbeatMS!
			// _ = esaclient.heartbeatMS

			lookupTable[*workUnit.req.ID] = workUnit.respChan

			reqMsgChan <- workUnit.req
		}
	}
}

// reader is responsible for reading all incoming messages and sending the corresponding objects down the channel
func reader(esaclient *ESAClient, respMsgChan chan<- ResponseMessage, stopChan <-chan bool) {
	// Create a buffer for incoming messages
	var buf []byte = make([]byte, esaclient.readerBufferSize)
	indiceStart := 0
	indiceStop := 0

	for {
		// check stopChan without blocking, if set then exit function
		select {
		case _, ok := <-stopChan:
			if !ok {
				close(respMsgChan)
				return
			}
		default:
		}

		// Call Read with timeout
		esaclient.conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		n, err := esaclient.conn.Read(buf[indiceStop : len(buf)-1])

		if esaclient.metricsFlag == 1 {
			if n != 0 {
				esaclient.readCounter.Add(float64(n))
			}
		}

		log.Log(globals.Logger, log.DEBUG, fmt.Sprintf("read %d bytes from connection", n), nil)

		if err1, ok := err.(*net.OpError); ok {
			// If timeout, continue
			if err1.Timeout() {
				log.Log(globals.Logger, log.DEBUG, "timeout occured while waiting to read", nil)
				continue
			}
			log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
			// Before continue, prob need to do something with the potential stuff that is in the buffer!
			// TODO: get the extra data read and update indices!

			// Inform the main program that connection was closed!
			return
		} else if err == io.EOF {
			// Connection was disconnected!
			log.Log(globals.Logger, log.ERROR, "connection closed on the server side", nil)

			// Inform the main program that connection was closed!
			return
		} else if err != nil {
			log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
			// Before continue, prob need to do something with the potential stuff that is in the buffer!
			// TODO: get the extra data read and update indices!

			// Inform the main program that connection was closed!
			return
		}

		fields := log.Fields{"message": fmt.Sprintf("%s", buf[indiceStart:n]), "type": "reader-data"}
		log.Log(globals.Logger, log.TRACE, "reader received message", fields)

		indiceStop += n
		progress := true

		// Iterate through the various potential messages (delim \r\n) we might have
		for progress {
			progress = false

			for i := indiceStart; i < indiceStop; i++ {
				if buf[i] == '\n' {
					if i-indiceStart == 0 || i-indiceStart == 1 {
						// discard
						indiceStart = i + 1
						progress = true
						break
					} else if buf[i-1] != '\r' {
						// discard
						indiceStart = i + 1
						progress = true
						break
					} else {
						// Log this into its own message field, and define a "type" field that specifies this is message related (like zap)
						log.Log(globals.Logger, log.DEBUG, fmt.Sprintf("message ready to be assembled: %s", buf[indiceStart:i-1]), nil)

						// Unmarshal into an object
						respMsg := ResponseMessage{}
						err = json.Unmarshal(buf[indiceStart:i-1], &respMsg)
						if err != nil {
							log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
						} else {
							// Push the newly created object down the channel
							esaclient.heartbeatUpdateChan <- 1
							respMsgChan <- respMsg
						}

						indiceStart = i + 1
						progress = true
						break
					}
				}
			}
		}

		// Shift content to the beginning of the buffer
		if indiceStart != 0 {
			if indiceStop == indiceStart {
				indiceStart = 0
				indiceStop = 0
			} else {
				copy(buf[:], buf[indiceStart:indiceStop])
				indiceStop -= indiceStart
				indiceStart = 0
			}
		}
	}
}

func writer(esaclient *ESAClient, reqMsgChan <-chan RequestMessage, stopInformChan chan<- bool) {
	for {
		select {
		case reqMsg, ok := <-reqMsgChan:
			if !ok {
				close(stopInformChan)
				return
			}

			// log.Log(globals.Logger, log.DEBUG, fmt.Sprintf("struct received on writer: %+v", reqMsg), nil)

			// Marhsal Request
			bytes, err := json.Marshal(reqMsg)
			if err != nil {
				log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
				continue
			}

			fields := log.Fields{"message": fmt.Sprintf("%s", bytes), "type": "writer-data"}
			log.Log(globals.Logger, log.TRACE, "writer sending message", fields)

			esaclient.conn.SetWriteDeadline(time.Now().Add(timeoutDuration))

			// TODO: Write function might not write all bytes to kernel!
			n, err := esaclient.conn.Write(append(bytes[:], []byte{'\r', '\n'}...))

			log.Log(globals.Logger, log.DEBUG, fmt.Sprintf("write %d bytes to connection", n), nil)

			if err1, ok := err.(*net.OpError); ok {
				// If timeout, continue
				if err1.Timeout() {
					log.Log(globals.Logger, log.DEBUG, "timeout occured while waiting to write", nil)
					continue
				}
				log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!

				// Inform the main program that connection was closed!
				return
			} else if err == io.EOF {
				// Connection was disconnected!
				log.Log(globals.Logger, log.ERROR, "connection closed on the server side", nil)

				// Inform the main program that connection was closed!
				return
			} else if err != nil {
				// If EOF, connection was closed!!
				log.Log(globals.Logger, log.ERROR, fmt.Sprintf("error: type [%T] - %+[1]v", err), nil)
				// Before continue, prob need to do something with the potential stuff that is in the buffer!
				// TODO: get the extra data read and update indices!

				// Inform the main program that connection was closed!
				return
			}
		}
	}
}

func connTracker(esaclient *ESAClient, stopChan <-chan bool, stopInformChan chan<- bool) {
	// Wait here until Auth has been done
	select {
	case _, ok := <-stopChan:
		if !ok {
			close(stopInformChan)
			return
		}
	case <-esaclient.authSuccessChan:
		// After Auth was successful then proceed
	}

	// Before calling MarketSubscribe or OrderSubscribe there won't be any heartbeats, but that's fine
	// because calls to the Heartbeat function must still succeed
	for {
		select {
		case _, ok := <-stopChan:
			if !ok {
				close(stopInformChan)
				return
			}
		case <-esaclient.heartbeatUpdateChan:
			// Got update
		case <-time.After(time.Duration(esaclient.heartbeatMultiplier*esaclient.heartbeatMS/100) * time.Millisecond):
			// Call heartbeat function
			fields := log.Fields{"source": "connTracker"}
			log.Log(globals.Logger, log.DEBUG, "time's up, sending heartbeat to check connection status", fields)
			sm, err := esaclient.Heartbeat()
			if err != nil {
				log.Log(globals.Logger, log.ERROR, "error while doing conn tracking", fields)

				// Inform the main program that connection is no longer good!

			} else if sm.StatusCode != StatusCode_Success {
				fields = log.Fields{"source": "connTracker", "StatusMessage": sm}
				log.Log(globals.Logger, log.ERROR, "heartbeat was not successfull! show status message", fields)

				// Inform the main program that connection is no longer good!
			} else {
				log.Log(globals.Logger, log.DEBUG, "got heartbeat response back, connection is good", fields)
			}
		}
	}
}
