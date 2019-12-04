package betting

type ContainerListMarketCatalogue struct {
	Filter           MarketFilter       `json:"filter"`
	MarketProjection []MarketProjection `json:"marketProjection,omitempty"`
	Sort             *MarketSort        `json:"sort,omitempty"`
	MaxResults       uint               `json:"maxResults"`
	Locale           *string            `json:"locale,omitempty"`
}

type ContainerListMarketBook struct {
	MarketIDs []string `json:"marketIds"`
}

type ContainerPlaceOrder struct {
	MarketID     string             `json:"marketId"`
	Instructions []PlaceInstruction `json:"instructions"`
}

type ContainerReplaceOrder struct {
	MarketID     string               `json:"marketId"`
	Instructions []ReplaceInstruction `json:"instructions"`
}

type ContainerCancelOrder struct {
	MarketID     string              `json:"marketId"`
	Instructions []CancelInstruction `json:"instructions"`
}
