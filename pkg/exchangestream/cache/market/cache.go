package market

type MarketCache struct {
	Clk          string
	InitialClk   string
	PublishTime  *int
	MarketID     string
	TradedVolume *float64
	// MarketDefinition *MarketDefinition
	Runners map[int64]RunnerCache // key is the RunnerID
}

func NewMarketCache() MarketCache {
	mc := MarketCache{}

	return mc
}

func (mc *MarketCache) Update() {

}

type RunnerCache struct {
	SelectionID                int64
	LastTradedPrice            *float64
	TradedVolume               *float64
	StartingPriceNear          *float64
	StartingPriceFar           *float64
	Traded                     []PriceStep
	AvailableToBack            []PriceStep
	AvailableToLay             []PriceStep
	StartingPriceBack          []PriceStep
	StartingPriceLay           []PriceStep
	BestAvailableToBack        []PriceStep
	BestAvailableToLay         []PriceStep
	BestDisplayAvailableToBack []PriceStep
	BestDisplayAvailableToLay  []PriceStep
}

type PriceStep struct {
	Position float64
	Price    float64
	Size     float64
}
