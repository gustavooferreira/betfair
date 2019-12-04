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
	// Event        string          `json:"event"`
}

type RunnerCatalog struct {
	SelectionID  uint              `json:"selectionId"`
	RunnerName   string            `json:"runnerName"`
	Handicap     float64           `json:"handicap"`
	SortPriority uint              `json:"sortPriority"`
	Metadata     map[string]string `json:"metadata"`
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
	MarketID string   `json:"marketId"`
	Runners  []Runner `json:"runners"`
}

type Runner struct {
	SelectionID     uint         `json:"selectionId"`
	Status          RunnerStatus `json:"status"`
	LastPriceTraded float64      `json:lastPriceTraded`
}

type PlaceInstruction struct {
	OrderType   OrderType  `json:"orderType"`
	SelectionID Side       `json:"selectionId"`
	Side        string     `json:"side"`
	LimitOrder  LimitOrder `json:"limitOrder"`
}

type LimitOrder struct {
	Size            string          `json:"size"`
	Price           string          `json:"price"`
	PersistenceType PersistenceType `json:"persistenceType"`
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
	Instruction         PlaceInstruction            `json:"instruction"`
	BetID               *string                     `json:"betId"`
	AveragePriceMatched *float64                    `json:"averagePriceMatched"`
	SizeMatched         *float64                    `json:"sizeMatched"`
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
