package exchangestream

type ConnectionMessage struct {
	ConnectionID string `json:"connectionId"`
}

type AuthenticationMessage struct {
	AppKey       string `json:"appKey"`
	SessionToken string `json:"session"`
}

type StatusMessage struct {
	ErrorMessage     string     `json:"errorMessage,omitempty"`
	ErrorCode        ErrorCode  `json:"errorCode,omitempty"`
	ConnectionID     string     `json:"connectionId"`
	ConnectionClosed *bool      `json:"connectionClosed"`
	StatusCode       StatusCode `json:"statusCode"`
}

type OrderSubscriptionMessage struct {
	SegmentationEnabled bool         `json:"segmentationEnabled"`
	OrderFilter         *OrderFilter `json:"orderFilter,omitempty"`
	Clk                 string       `json:"clk,omitempty"`
	HeartbeatMs         uint         `json:"heartbeatMs,omitempty"`
	InitialClk          string       `json:"initialClk,omitempty"`
	ConflateMs          uint         `json:"conflateMs"`
}

type OrderFilter struct {
	IncludeOverallPosition        *bool    `json:"includeOverallPosition,omitempty"`
	CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`
	PartitionMatchedByStrategyRef *bool    `json:"partitionMatchedByStrategyRef,omitempty"`
}

type MarketSubscriptionMessage struct {
	SegmentationEnabled bool             `json:"segmentationEnabled"`
	Clk                 string           `json:"clk,omitempty"`
	HeartbeatMs         uint             `json:"heartbeatMs,omitempty"`
	InitialClk          string           `json:"initialClk,omitempty"`
	MarketFilter        MarketFilter     `json:"marketFilter"`
	ConflateMs          uint             `json:"conflateMs"`
	MarketDataFilter    MarketDataFilter `json:"marketDataFilter"`
}

type MarketFilter struct {
	CountryCodes      []string      `json:"countryCodes,omitempty"`
	BettingTypes      []BettingType `json:"bettingTypes,omitempty"`
	TurnInPlayEnabled *bool         `json:"turnInPlayEnabled,omitempty"`
	MarketTypes       []string      `json:"marketTypes,omitempty"`
	Venues            []string      `json:"venues,omitempty"`
	MarketIDs         []string      `json:"marketIds,omitempty"`
	EventTypeIDs      []string      `json:"eventTypeIds,omitempty"`
	EventIDs          []string      `json:"eventIds,omitempty"`
	BSPMarket         *bool         `json:"bspMarket,omitempty"`
	RaceTypes         []string      `json:"raceTypes,omitempty"`
}

type MarketDataFilter struct {
	LadderLevels uint        `json:"ladderLevels,omitempty"`
	Fields       []PriceData `json:"fields,omitempty"`
}

type MarketChangeMessage struct {
	ChangeType    *ChangeType    `json:"ct,omitempty"`
	Clk           string         `json:"clk"`
	HeartbeatMs   uint           `json:"heartbeatMs"`
	PublishTime   uint           `json:"pt"`
	InitialClk    string         `json:"initialClk"`
	MarketChanges []MarketChange `json:"mc,omitempty"`
	ConflateMs    uint           `json:"conflateMs"`
	SegmentType   *SegmentType   `json:"segmentType,omitempty"`
	Status        *uint          `json:"status,omitempty"`
}

type MarketChange struct {
	RunnerChanges    []RunnerChange    `json:"rc,omitempty"`
	Image            *bool             `json:"img,omitempty"`
	TotalVolume      *float64          `json:"tv"`
	Conflated        *bool             `json:"con,omitempty"`
	MarketDefinition *MarketDefinition `json:"marketDefinition,omitempty"`
	ID               string            `json:"id"`
}

type MarketDefinition struct {
	Venue                 string                `json:"venue"`
	RaceType              string                `json:"raceType"`
	SettledTime           string                `json:"settledTime"` // TODO: time.Time!
	Timezone              string                `json:"timezone"`
	EachWayDivisor        float64               `json:"eachWayDivisor"`
	Regulators            []string              `json:"regulators"`
	MarketType            string                `json:"marketType"`
	MarketBaseRate        float64               `json:"marketBaseRate"`
	NumberOfWinners       uint                  `json:"numberOfWinners"`
	CountryCode           string                `json:"countryCode"`
	LineMaxUnit           float64               `json:"lineMaxUnit"`
	InPlay                *bool                 `json:"inPlay,omitempty"`
	BetDelay              uint                  `json:"betDelay"`
	BSPMarket             *bool                 `json:"bspMarket,omitempty"`
	BettingType           BettingType           `json:"bettingType"`
	NumberOfActiveRunners uint                  `json:"numberOfActiveRunners"`
	LineMinUnit           float64               `json:"lineMinUnit"`
	EventID               string                `json:"eventId"`
	CrossMatching         *bool                 `json:"crossMatching,omitempty"`
	RunnersVoidable       *bool                 `json:"runnersVoidable,omitempty"`
	TurnInPlayEnabled     *bool                 `json:"turnInPlayEnabled,omitempty"`
	PriceLadderDefinition PriceLadderDefinition `json:"priceLadderDefinition"`
	KeyLineDefinition     KeyLineDefinition     `json:"keyLineDefinition"`
	SuspendTime           string                `json:"suspendTime"` // TODO: time.Time!
	DiscountAllowed       *bool                 `json:"discountAllowed,omitempty"`
	PersistenceEnabled    *bool                 `json:"persistenceEnabled,omitempty"`
	Runners               []RunnerDefinition    `json:"runners"`
	Version               uint                  `json:"version"`
	EventTypeID           string                `json:"eventTypeId"`
	Complete              *bool                 `json:"complete,omitempty"`
	OpenDate              string                `json:"openDate"`   // TODO: time.Time!
	MarketTime            string                `json:"marketTime"` // TODO: time.Time!
	BSPReconciled         *bool                 `json:"bspReconciled,omitempty"`
	LineInterval          float64               `json:"lineInterval"`
	Status                RaceStatus            `json:"status"`
}

type RunnerDefinition struct {
	SortPriority     uint       `json:"sortPriority"`
	RemovalDate      string     `json:"removalDate"` // TODO: time.Time!
	ID               uint       `json:"id"`
	Handicap         *float64   `json:"hc,omitempty"`
	AdjustmentFactor float64    `json:"adjustmentFactor"`
	BSP              float64    `json:"bsp"`
	Status           RaceStatus `json:"status"`
}

type PriceLadderDefinition struct {
	Type PriceLadderType `json:"type"`
}

type KeyLineDefinition struct {
	Kl []KeyLineSelection `json:"kl"`
}

type KeyLineSelection struct {
	ID       uint    `json:"id"`
	Handicap float64 `json:"hc"`
}

type RunnerChange struct {
	TotalVolume *float64    `json:"tv,omitempty"`
	BATB        [][]float64 `json:"batb,omitempty"`
	SPB         [][]float64 `json:"spb",omitempty`
	BDATL       [][]float64 `json:"bdatl",omitempty`
	TRD         [][]float64 `json:"trd",omitempty`
	SPF         *float64    `json:"spf",omitempty`
	LTP         *float64    `json:"ltp",omitempty`
	ATB         [][]float64 `json:"atb",omitempty`
	SPL         [][]float64 `json:"spl",omitempty`
	SPN         *float64    `json:"spn",omitempty`
	ATL         [][]float64 `json:"atl",omitempty`
	BATL        [][]float64 `json:"batl",omitempty`
	ID          uint        `json:"id",omitempty`
	Handicap    *float64    `json:"hc",omitempty`
	BDATB       [][]float64 `json:"bdatb",omitempty`
}

type OrderChangeMessage struct {
	ChangeType         *ChangeType         `json:"ct,omitempty"`
	Clk                string              `json:"clk"`
	HeartbeatMs        uint                `json:"heartbeatMs"`
	PublishTime        uint                `json:"pt"`
	OrderMarketChanges []OrderMarketChange `json:"oc,omitempty"`
	InitialClk         string              `json:"initialClk"`
	ConflateMs         uint                `json:"conflateMs"`
	SegmentType        *SegmentType        `json:"segmentType,omitempty"`
	Status             *uint               `json:"status,omitempty"`
}

type OrderMarketChange struct {
	AccountID    uint                `json:"accountId"`
	OrderChanges []OrderRunnerChange `json:"orc,omitempty"`
	Closed       *bool               `json:"closed"`
	ID           string              `json:"id"`
}

type OrderRunnerChange struct {
	MatchedBacks    [][]float64                    `json:"mb,omitempty"`
	StrategyMatches map[string]StrategyMatchChange `json:"smc",omitempty`
	UnmatchedOrders []Order                        `json:"uo",omitempty`
	ID              uint                           `json:"id"`
	Handicap        float64                        `json:"hc",omitempty`
	FullImage       *bool                          `json:"fullImage",omitempty`
	MatchedLays     [][]float64                    `json:"ml",omitempty`
}

type Order struct {
	Side                OrderSide       `json:"side"`
	SizeVoided          float64         `json:"sv"`
	PersistenceType     PersistenceType `json:"pt"`
	OrderType           OrderType       `json:"ot"`
	Price               float64         `json:"p"`
	SizeCancelled       float64         `json:"sc"`
	RegulatorCode       string          `json:"rc"`
	Size                float64         `json:"s"`
	PlacedDate          uint            `json:"pd"` //TODO: time.Time!
	RegulatorAuthCode   string          `json:"rac"`
	MatchedDate         *uint           `json:"md,omitempty"` //TODO: time.Time!
	LapsedDate          *uint           `json:"ld,omitempty"` //TODO: time.Time!
	SizeLapsed          float64         `json:"sl"`
	AveragePriceMatched *float64        `json:"avp,omitempty"`
	SizeMatched         float64         `json:"sm"`
	OrderReference      *string         `json:"rfo,omitempty"`
	ID                  string          `json:"id"`
	BSP                 *float64        `json:"bsp,omitempty"`
	StrategyReference   *string         `json:"rfs,omitempty"`
	Status              OrderStatus     `json:"status"`
	SizeRemaining       float64         `json:"sr"`
}

type StrategyMatchChange struct {
	MatchedBacks [][]float64 `json:"mb,omitempty"`
	MatchedLays  [][]float64 `json:"ml,omitempty"`
}
