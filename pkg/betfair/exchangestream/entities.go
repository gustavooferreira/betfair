package exchangestream

import (
	"encoding/json"
	"errors"
)

type RequestMessage struct {
	Op string `json:"op"`
	ID uint32 `json:"id"`
}

type ResponseMessage struct {
	Op                  string `json:"op"`
	ID                  uint32 `json:"id"`
	StatusMessage       *StatusMessage
	MarketChangeMessage *MarketChangeMessage
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

	rm.StatusMessage = nil
	rm.MarketChangeMessage = nil

	if rm.Op == "status" {
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

type ConnectionMessage struct {
	Op           string `json:"op"`
	ConnectionID string `json:"connectionId"`
}

type AuthenticationMessage struct {
	RequestMessage
	AppKey       string `json:"appKey"`
	SessionToken string `json:"session"`
}

type StatusMessage struct {
	StatusCode       string `json:"statusCode"`
	ConnectionClosed *bool  `json:"connectionClosed"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorCode        string `json:"errorCode"`
}

type SubscriptionMessage struct {
	RequestMessage
	SegmentationEnabled *bool            `json:"segmentationEnabled,omitempty"`
	Clk                 string           `json:"clk,omitempty"`
	InitialClk          string           `json:"initialClk,omitempty"`
	HeartbeatMs         uint             `json:"heartbeatMs,omitempty"`
	ConflateMs          uint             `json:"conflateMs,omitempty"`
	MarketFilter        MarketFilter     `json:"marketFilter,omitempty"`
	MarketDataFilter    MarketDataFilter `json:"marketDataFilter,omitempty"`
}

type MarketFilter struct {
	CountryCodes      []string `json:"countryCodes,omitempty"`
	BettingTypes      []string `json:"bettingTypes,omitempty"`
	TurnInPlayEnabled *bool    `json:"turnInPlayEnabled,omitempty"`
	MarketTypes       []string `json:"marketTypes,omitempty"`
	Venues            []string `json:"venues,omitempty"`
	MarketIDs         []string `json:"marketIds,omitempty"`
	EventTypeIDs      []string `json:"eventTypeIds,omitempty"`
	EventIDs          []string `json:"eventIds,omitempty"`
	BSPMarket         *bool    `json:"bspMarket,omitempty"`
	RaceTypes         []string `json:"raceTypes,omitempty"`
}

type MarketDataFilter struct {
	LadderLevels uint     `json:"ladderLevels,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

type MarketChangeMessage struct {
	Ct          string         `json:"ct"`
	Clk         string         `json:"clk"`
	InitialClk  string         `json:"initialClk"`
	Pt          uint           `json:"pt"`
	HeartbeatMs uint           `json:"heartbeatMs"`
	ConflateMs  uint           `json:"conflateMs"`
	SegmentType string         `json:"segmentType"`
	Status      uint           `json:"status"`
	Mc          []MarketChange `json:"mc"`
}

type MarketChange struct {
	Img              *bool            `json:"img"`
	Tv               float64          `json:"tv"`
	Con              *bool            `json:"con"`
	ID               string           `json:"id"`
	Rc               []RunnerChange   `json:"rc"`
	MarketDefinition MarketDefinition `json:"marketDefinition"`
}

type MarketDefinition struct {
	Venue                 string                `json:"venue"`
	RaceType              string                `json:"raceType"`
	SettledTime           string                `json:"settledTime"`
	Timezone              string                `json:"timezone"`
	EachWayDivisor        float64               `json:"eachWayDivisor"`
	Regulators            []string              `json:"regulators"`
	MarketType            string                `json:"marketType"`
	MarketBaseRate        float64               `json:"marketBaseRate"`
	NumberOfWinners       uint                  `json:"numberOfWinners"`
	CountryCode           string                `json:"countryCode"`
	LineMaxUnit           float64               `json:"lineMaxUnit"`
	InPlay                *bool                 `json:"inPlay"`
	BetDelay              uint                  `json:"betDelay"`
	BSPMarket             *bool                 `json:"bspMarket"`
	BettingType           string                `json:"bettingType"`
	NumberOfActiveRunners uint                  `json:"numberOfActiveRunners"`
	LineMinUnit           float64               `json:"lineMinUnit"`
	EventID               string                `json:"eventId"`
	CrossMatching         *bool                 `json:"crossMatching"`
	RunnersVoidable       *bool                 `json:"runnersVoidable"`
	TurnInPlayEnabled     *bool                 `json:"turnInPlayEnabled"`
	PriceLadderDefinition PriceLadderDefinition `json:"priceLadderDefinition"`
	KeyLineDefinition     KeyLineDefinition     `json:"keyLineDefinition"`
	SuspendTime           string                `json:"suspendTime"`
	DiscountAllowed       *bool                 `json:"discountAllowed"`
	PersistenceEnabled    *bool                 `json:"persistenceEnabled"`
	Runners               []RunnerDefinition    `json:"runners"`
	Version               uint                  `json:"version"`
	EventTypeID           string                `json:"eventTypeId"`
	Complete              *bool                 `json:"complete"`
	OpenDate              string                `json:"openDate"`
	MarketTime            string                `json:"marketTime"`
	BSPReconciled         *bool                 `json:"bspReconciled"`
	LineInterval          float64               `json:"lineInterval"`
	Status                string                `json:"status"`
}

type RunnerDefinition struct {
	SortPriority     uint    `json:"sortPriority"`
	RemovalDate      string  `json:"removalDate"`
	ID               uint    `json:"id"`
	Hc               float64 `json:"hc"`
	AdjustmentFactor float64 `json:"adjustmentFactor"`
	BSP              float64 `json:"bsp"`
	Status           string  `json:"status"`
}

type PriceLadderDefinition struct {
	Type string `json:"type"`
}

type KeyLineDefinition struct {
	Kl []KeyLineSelection `json:"kl"`
}

type KeyLineSelection struct {
	Id uint    `json:"id"`
	Hc float64 `json:"hc"`
}

type RunnerChange struct {
	Tv    float64     `json:"tv"`
	Batb  [][]float64 `json:"batb"`
	Spb   [][]float64 `json:"spb"`
	Bdatl [][]float64 `json:"bdatl"`
	Trd   [][]float64 `json:"trd"`
	Atb   [][]float64 `json:"atb"`
	Spl   [][]float64 `json:"spl"`
	Atl   [][]float64 `json:"atl"`
	Batl  [][]float64 `json:"batl"`
	Bdatb [][]float64 `json:"bdatb"`
	Spf   float64     `json:"spf"`
	Ltp   float64     `json:"ltp"`
	Spn   float64     `json:"spn"`
	ID    uint        `json:"id"`
	Hc    float64     `json:"hc"`
}
