package betting

import (
	"time"
)

type MarketFilter struct {
	TextQuery          string              `json:"textQuery,omitempty"`
	ExchangeIDs        []string            `json:"exchangeIds,omitempty"` //NOTE: Deprecated
	EventTypeIDs       []string            `json:"eventTypeIds,omitempty"`
	EventIds           []string            `json:"eventIds,omitempty"`
	CompetitionISs     []string            `json:"competitionIds,omitempty"`
	MarketIds          []string            `json:"marketIds,omitempty"`
	Venues             []string            `json:"venues,omitempty"`
	BSPOnly            *bool               `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  *bool               `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         *bool               `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []MarketBettingType `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string            `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string            `json:"marketTypeCodes,omitempty"`
	MarketStartTime    *TimeRange          `json:"marketStartTime,omitempty"`
	WithOrders         []OrderStatus       `json:"withOrders,omitempty"`
	RaceTypes          []string            `json:"raceTypes,omitempty"`
}

type TimeRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type MarketCatalogue struct {
	MarketID        string     `json:"marketId"`
	MarketName      string     `json:"marketName"`
	MarketStartTime *time.Time `json:"marketStartTime"`
	// Description     string          `json:"description"`
	TotalMatched float64         `json:"totalMatched"`
	Runners      []RunnerCatalog `json:"runners"`
	// EventType    string          `json:"eventType"`
	// Competition  string          `json:"competition"`
	Event Event `json:"event"`
}

type RunnerCatalog struct {
	SelectionID  uint              `json:"selectionId"`
	RunnerName   string            `json:"runnerName"`
	Handicap     float64           `json:"handicap"`
	SortPriority uint              `json:"sortPriority"`
	Metadata     map[string]string `json:"metadata"`
}

type Event struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CountryCode string     `json:"countryCode"`
	Timezone    string     `json:"timezone"`
	Venue       string     `json:"venue"`
	OpenDate    *time.Time `json:"openDate"`
}

type BetfairAPIError struct {
	Detail      BetfairDetailError `json:"detail"`
	FaultCode   string             `json:"faultCode"`
	FaultString string             `json:"faultstring"`
}

type BetfairDetailError struct {
	APINGException APINGException `json:"APINGException"`
	ExceptionName  string         `json:"exceptionname"`
}

type APINGException struct {
	ErrorCode    APINGExceptionCode `json:"errorCode"`
	ErrorDetails string             `json:"errorDetails"`
	RequestUUID  string             `json:"requestUUID"`
}

type MarketBook struct {
	MarketID              string         `json:"marketId"`
	Status                MarketStatus   `json:"status"`
	InPlay                bool           `json:"inplay"`
	NumberOfWinners       uint           `json:"numberOfWinners"`
	NumberOfRunners       uint           `json:"numberOfRunners"`
	NumberOfActiveRunners uint           `json:"numberOfActiveRunners"`
	LastMatchTime         time.Time      `json:"lastMatchTime"`
	TotalMatched          float64        `json:"totalMatched"`
	TotalAvailable        float64        `json:"totalAvailable"`
	Version               uint           `json:"version"`
	Runners               []Runner       `json:"runners"`
	Ex                    ExchangePrices `json:"ex"`
}

type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack"`
	AvailableToLay  []PriceSize `json:"availableToLay"`
	TradedVolume    []PriceSize `json:"tradedVolume"`
}

type PriceSize struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Runner struct {
	SelectionID     uint         `json:"selectionId"`
	Status          RunnerStatus `json:"status"`
	LastPriceTraded float64      `json:lastPriceTraded`
	TotalMatched    float64      `json:totalMatched`
}

type PlaceInstruction struct {
	OrderType        OrderType  `json:"orderType"`
	SelectionID      uint       `json:"selectionId"`
	Side             Side       `json:"side"`
	LimitOrder       LimitOrder `json:"limitOrder"`
	CustomerOrderRef string     `json:"customerOrderRef"`
}

type LimitOrder struct {
	Size            string          `json:"size"`
	Price           string          `json:"price"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
}

type PlaceExecutionReport struct {
	Status             ExecutionReportStatus    `json:"status"`
	ErrorCode          ExecutionReportErrorCode `json:"errorCode"`
	MarketID           string                   `json:"marketId"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports"`
}

type PlaceInstructionReport struct {
	Status              InstructionReportStatus     `json:"status"`
	ErrorCode           *InstructionReportErrorCode `json:"errorCode"`
	OrderStatus         OrderStatus                 `json:"orderStatus"`
	Instruction         PlaceInstructionR           `json:"instruction"`
	BetID               *string                     `json:"betId"`
	AveragePriceMatched *float64                    `json:"averagePriceMatched"`
	SizeMatched         *float64                    `json:"sizeMatched"`
}

type PlaceInstructionR struct {
	OrderType   OrderType   `json:"orderType"`
	SelectionID uint        `json:"selectionId"`
	Side        Side        `json:"side"`
	LimitOrder  LimitOrderR `json:"limitOrder"`
}

type LimitOrderR struct {
	Size            float64         `json:"size"`
	Price           float64         `json:"price"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
}

type ReplaceInstruction struct {
	BetID    string  `json:"betId"`
	NewPrice float64 `json:"newPrice"`
}

type ReplaceExecutionReport struct {
	Status             ExecutionReportStatus      `json:"status"`
	ErrorCode          ExecutionReportErrorCode   `json:"errorCode"`
	MarketID           string                     `json:"marketId"`
	InstructionReports []ReplaceInstructionReport `json:"instructionReports"`
}

type ReplaceInstructionReport struct {
	Status    InstructionReportStatus     `json:"status"`
	ErrorCode *InstructionReportErrorCode `json:"errorCode"`
}

type CancelInstruction struct {
	BetID         string   `json:"betId"`
	SizeReduction *float64 `json:"sizeReduction"`
}

type CancelExecutionReport struct {
	Status             ExecutionReportStatus     `json:"status"`
	ErrorCode          ExecutionReportErrorCode  `json:"errorCode"`
	MarketID           string                    `json:"marketId"`
	InstructionReports []CancelInstructionReport `json:"instructionReports"`
}

type CancelInstructionReport struct {
	Status        InstructionReportStatus     `json:"status"`
	ErrorCode     *InstructionReportErrorCode `json:"errorCode"`
	Instruction   CancelInstruction           `json:"instruction"`
	SizeCancelled float64                     `json:"sizeCancelled"`
}

type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

type ClearedOrderSummary struct {
	MarketID            string  `json:"marketId"`
	SelectionID         uint    `json:"SelectionId"`
	Side                Side    `json:"side"`
	BetOutcome          string  `json:"betOutcome"`
	Commission          float64 `json:"commission"`
	PriceMatched        float64 `json:"priceMatched"`
	Profit              float64 `json:"profit"`
	CustomerStrategyRef string  `json:"customerStrategyRef"`
}

type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

type CurrentOrderSummary struct {
	BetID  string      `json:"betId"`
	Status OrderStatus `json:"status"`
}
