package exchangestream

import (
	"encoding/json"
	"errors"
)

type RequestMessage struct {
	Op                        string `json:"op"`
	ID                        uint32 `json:"id"`
	AuthenticationMessage     *AuthenticationMessage
	MarketSubscriptionMessage *MarketSubscriptionMessage
	OrderSubscriptionMessage  *OrderSubscriptionMessage
}

// MarshalJSON marshals RequestMessage struct
func (rm RequestMessage) MarshalJSON() ([]byte, error) {
	if rm.Op == "heartbeat" {
		return json.Marshal(&struct {
			Op string `json:"op"`
			ID uint32 `json:"id"`
		}{
			Op: rm.Op,
			ID: rm.ID,
		})
	} else if rm.Op == "authentication" && rm.AuthenticationMessage != nil {
		return json.Marshal(&struct {
			Op           string `json:"op"`
			ID           uint32 `json:"id"`
			AppKey       string `json:"appKey"`
			SessionToken string `json:"session"`
		}{
			Op:           rm.Op,
			ID:           rm.ID,
			AppKey:       rm.AuthenticationMessage.AppKey,
			SessionToken: rm.AuthenticationMessage.SessionToken,
		})
	} else if rm.Op == "marketSubscription" && rm.MarketSubscriptionMessage != nil {
		return json.Marshal(&struct {
			Op string `json:"op"`
			ID uint32 `json:"id"`
		}{
			Op: rm.Op,
			ID: rm.ID,
		})
	} else if rm.Op == "orderSubscription" && rm.OrderSubscriptionMessage != nil {
		return json.Marshal(&struct {
			Op                  string      `json:"op"`
			ID                  uint32      `json:"id"`
			SegmentationEnabled bool        `json:"segmentationEnabled"`
			OrderFilter         OrderFilter `json:"orderFilter"`
			Clk                 string      `json:"clk,omitempty"`
			HeartbeatMs         uint        `json:"heartbeatMs,omitempty"`
			InitialClk          string      `json:"initialClk,omitempty"`
			ConflateMs          uint        `json:"conflateMs,omitempty"`
		}{
			Op:                  rm.Op,
			ID:                  rm.ID,
			SegmentationEnabled: rm.OrderSubscriptionMessage.SegmentationEnabled,
			OrderFilter:         rm.OrderSubscriptionMessage.OrderFilter,
			Clk:                 rm.OrderSubscriptionMessage.Clk,
			HeartbeatMs:         rm.OrderSubscriptionMessage.HeartbeatMs,
			InitialClk:          rm.OrderSubscriptionMessage.InitialClk,
			ConflateMs:          rm.OrderSubscriptionMessage.ConflateMs,
		})
	}

	return json.Marshal(nil)
}

type ResponseMessage struct {
	Op                  string `json:"op"`
	ID                  uint32 `json:"id"`
	ConnectionMessage   *ConnectionMessage
	StatusMessage       *StatusMessage
	MarketChangeMessage *MarketChangeMessage
	OrderChangeMessage  *OrderChangeMessage
}

// UnmarshalJSON unmarshals ResponseMessage struct
func (rm *ResponseMessage) UnmarshalJSON(data []byte) error {
	temp := struct {
		Op string `json:"op"`
		ID uint32 `json:"id"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	rm.Op = temp.Op
	rm.ID = temp.ID

	if rm.Op == "connection" {
		var connectionMessage ConnectionMessage
		if err := json.Unmarshal(data, &connectionMessage); err != nil {
			return err
		}
		rm.ConnectionMessage = &connectionMessage
	} else if rm.Op == "status" {
		var statusMessage StatusMessage
		if err := json.Unmarshal(data, &statusMessage); err != nil {
			return err
		}
		rm.StatusMessage = &statusMessage
	} else if rm.Op == "mcm" {
		var marketChangeMessage MarketChangeMessage
		if err := json.Unmarshal(data, &marketChangeMessage); err != nil {
			return err
		}
		rm.MarketChangeMessage = &marketChangeMessage
	} else {
		return errors.New("Invalid object value")
	}

	return nil
}
