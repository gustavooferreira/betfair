package exchangestream

type MarketChangeM struct {
	ID *uint32
	MarketChangeMessage
}

type OrderChangeM struct {
	ID *uint32
	OrderChangeMessage
}
