package exchangestream

// Consts
const serverHost string = "stream-api.betfair.com"
const serverPort string = "443"

type ConnManager struct {
	ESAClient

	// RWLock
	idCount uint32
}

func (connM *ConnManager) getNewID() uint32 {
	id := connM.idCount
	connM.idCount++
	return id
}
