package betting

type ContainerListMarketCatalogue struct {
	Filter           MarketFilter       `json:"filter"`
	MarketProjection []MarketProjection `json:"marketProjection,omitempty"`
	Sort             *MarketSort        `json:"sort,omitempty"`
	MaxResults       uint               `json:"maxResults"`
	Locale           *string            `json:"locale,omitempty"`
}

type ContainerListMarketBook struct {
	MarketIDs       []string        `json:"marketIds"`
	PriceProjection PriceProjection `json:"priceProjection"`
}

type ContainerPlaceOrders struct {
	MarketID     string             `json:"marketId"`
	Instructions []PlaceInstruction `json:"instructions"`
	// CustomerStrategyRef - Max of 15 characters
	CustomerStrategyRef string `json:"customerStrategyRef"`
}

type ContainerReplaceOrders struct {
	MarketID     string               `json:"marketId"`
	Instructions []ReplaceInstruction `json:"instructions"`
}

type ContainerCancelOrders struct {
	MarketID     string              `json:"marketId"`
	Instructions []CancelInstruction `json:"instructions"`
}

type ContainerListClearedOrders struct {
	BetStatus            BetStatus  `json:"betStatus"`
	GroupBy              *GroupBy   `json:"groupBy"`
	CustomerStrategyRefs []string   `json:"customerStrategyRefs,omitempty"`
	SettleDateRange      *TimeRange `json:"settleDateRange,omitempty"`
}

type ContainerListCurrentOrders struct {
	BetIDs []string `json:"betIds,omitempty"`
}
