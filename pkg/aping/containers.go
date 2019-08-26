package aping

type listMarketCatalogueReqContainer struct {
	Filter           MarketFilter        `json:"filter"`
	MarketProjection *[]MarketProjection `json:"marketProjection,omitempty"`
	Sort             *MarketSort         `json:"sort, omitempty"`
	MaxResults       uint                `json:"maxResults"`
}

type listMarketBookReqContainer struct {
	MarketIDs []string `json:"marketIds"`
}
