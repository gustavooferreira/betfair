package exchangestream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// ErrorCode ENUM

type ErrorCode int

const (
	ErrorCode_NoAppKey ErrorCode = iota + 1
	ErrorCode_InvalidAppKey
	ErrorCode_NoSession
	ErrorCode_InvalidSessionInformation
	ErrorCode_NotAuthorized
	ErrorCode_InvalidInput
	ErrorCode_InvalidClock
	ErrorCode_UnexpectedError
	ErrorCode_Timeout
	ErrorCode_SubscriptionLimitExceeded
	ErrorCode_InvalidRequest
	ErrorCode_ConnectionFailed
	ErrorCode_MaxConnectionLimitExceeded
)

func (ec ErrorCode) String() string {
	return errorCodeToString[ec]
}

var errorCodeToString = map[ErrorCode]string{
	ErrorCode_NoAppKey:                   "NO_APP_KEY",
	ErrorCode_InvalidAppKey:              "INVALID_APP_KEY",
	ErrorCode_NoSession:                  "NO_SESSION",
	ErrorCode_InvalidSessionInformation:  "INVALID_SESSION_INFORMATION",
	ErrorCode_NotAuthorized:              "NOT_AUTHORIZED",
	ErrorCode_InvalidInput:               "INVALID_INPUT",
	ErrorCode_InvalidClock:               "INVALID_CLOCK",
	ErrorCode_UnexpectedError:            "UNEXPECTED_ERROR",
	ErrorCode_Timeout:                    "TIMEOUT",
	ErrorCode_SubscriptionLimitExceeded:  "SUBSCRIPTION_LIMIT_EXCEEDED",
	ErrorCode_InvalidRequest:             "INVALID_REQUEST",
	ErrorCode_ConnectionFailed:           "CONNECTION_FAILED",
	ErrorCode_MaxConnectionLimitExceeded: "MAX_CONNECTION_LIMIT_EXCEEDED",
}

var errorCodeToEnum = map[string]ErrorCode{
	"NO_APP_KEY":                    ErrorCode_NoAppKey,
	"INVALID_APP_KEY":               ErrorCode_InvalidAppKey,
	"NO_SESSION":                    ErrorCode_NoSession,
	"INVALID_SESSION_INFORMATION":   ErrorCode_InvalidSessionInformation,
	"NOT_AUTHORIZED":                ErrorCode_NotAuthorized,
	"INVALID_INPUT":                 ErrorCode_InvalidInput,
	"INVALID_CLOCK":                 ErrorCode_InvalidClock,
	"UNEXPECTED_ERROR":              ErrorCode_UnexpectedError,
	"TIMEOUT":                       ErrorCode_Timeout,
	"SUBSCRIPTION_LIMIT_EXCEEDED":   ErrorCode_SubscriptionLimitExceeded,
	"INVALID_REQUEST":               ErrorCode_InvalidRequest,
	"CONNECTION_FAILED":             ErrorCode_ConnectionFailed,
	"MAX_CONNECTION_LIMIT_EXCEEDED": ErrorCode_MaxConnectionLimitExceeded,
}

// MarshalJSON marshals the enum as a quoted json string
func (ec ErrorCode) MarshalJSON() ([]byte, error) {
	elem, ok := errorCodeToString[ec]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ec *ErrorCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := errorCodeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ec = result
	return nil
}

// StatusCode ENUM

type StatusCode int

const (
	StatusCode_Success StatusCode = iota + 1
	StatusCode_Failure
)

func (sc StatusCode) String() string {
	return statusCodeToString[sc]
}

var statusCodeToString = map[StatusCode]string{
	StatusCode_Success: "SUCCESS",
	StatusCode_Failure: "FAILURE",
}

var statusCodeToEnum = map[string]StatusCode{
	"SUCCESS": StatusCode_Success,
	"FAILURE": StatusCode_Failure,
}

// MarshalJSON marshals the enum as a quoted json string
func (sc StatusCode) MarshalJSON() ([]byte, error) {
	elem, ok := statusCodeToString[sc]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (sc *StatusCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := statusCodeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*sc = result
	return nil
}

// PriceData ENUM

type PriceData int

const (
	PriceData_ExBestOffersDisp PriceData = iota + 1
	PriceData_ExBestOffers
	PriceData_ExAllOffers
	PriceData_ExTraded
	PriceData_ExTradedVol
	PriceData_ExLTP
	PriceData_ExMarketDef
	PriceData_SPTraded
	PriceData_SPProjected
)

func (pd PriceData) String() string {
	return priceDataToString[pd]
}

var priceDataToString = map[PriceData]string{
	PriceData_ExBestOffersDisp: "EX_BEST_OFFERS_DISP",
	PriceData_ExBestOffers:     "EX_BEST_OFFERS",
	PriceData_ExAllOffers:      "EX_ALL_OFFERS",
	PriceData_ExTraded:         "EX_TRADED",
	PriceData_ExTradedVol:      "EX_TRADED_VOL",
	PriceData_ExLTP:            "EX_LTP",
	PriceData_ExMarketDef:      "EX_MARKET_DEF",
	PriceData_SPTraded:         "SP_TRADED",
	PriceData_SPProjected:      "SP_PROJECTED",
}

var priceDataToEnum = map[string]PriceData{
	"EX_BEST_OFFERS_DISP": PriceData_ExBestOffersDisp,
	"EX_BEST_OFFERS":      PriceData_ExBestOffers,
	"EX_ALL_OFFERS":       PriceData_ExAllOffers,
	"EX_TRADED":           PriceData_ExTraded,
	"EX_TRADED_VOL":       PriceData_ExTradedVol,
	"EX_LTP":              PriceData_ExLTP,
	"EX_MARKET_DEF":       PriceData_ExMarketDef,
	"SP_TRADED":           PriceData_SPTraded,
	"SP_PROJECTED":        PriceData_SPProjected,
}

// MarshalJSON marshals the enum as a quoted json string
func (pd PriceData) MarshalJSON() ([]byte, error) {
	elem, ok := priceDataToString[pd]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (pd *PriceData) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := priceDataToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*pd = result
	return nil
}

// BettingType ENUM

type BettingType int

const (
	BettingType_Odds BettingType = iota + 1
	BettingType_Line
	BettingType_Range
	BettingType_AsianHandicapDoubleLine
	BettingType_AsianHandicapSingleLine
)

func (bt BettingType) String() string {
	return bettingTypeToString[bt]
}

var bettingTypeToString = map[BettingType]string{
	BettingType_Odds:                    "ODDS",
	BettingType_Line:                    "LINE",
	BettingType_Range:                   "RANGE",
	BettingType_AsianHandicapDoubleLine: "ASIAN_HANDICAP_DOUBLE_LINE",
	BettingType_AsianHandicapSingleLine: "ASIAN_HANDICAP_SINGLE_LINE",
}

var bettingTypeToEnum = map[string]BettingType{
	"ODDS":                       BettingType_Odds,
	"LINE":                       BettingType_Line,
	"RANGE":                      BettingType_Range,
	"ASIAN_HANDICAP_DOUBLE_LINE": BettingType_AsianHandicapDoubleLine,
	"ASIAN_HANDICAP_SINGLE_LINE": BettingType_AsianHandicapSingleLine,
}

// MarshalJSON marshals the enum as a quoted json string
func (bt BettingType) MarshalJSON() ([]byte, error) {
	elem, ok := bettingTypeToString[bt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (bt *BettingType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := bettingTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*bt = result
	return nil
}

// ChangeType ENUM

type ChangeType int

const (
	ChangeType_SubImage ChangeType = iota + 1
	ChangeType_ResubDelta
	ChangeType_Heartbeat
)

func (ct ChangeType) String() string {
	return changeTypeToString[ct]
}

var changeTypeToString = map[ChangeType]string{
	ChangeType_SubImage:   "SUB_IMAGE",
	ChangeType_ResubDelta: "RESUB_DELTA",
	ChangeType_Heartbeat:  "HEARTBEAT",
}

var changeTypeToEnum = map[string]ChangeType{
	"SUB_IMAGE":   ChangeType_SubImage,
	"RESUB_DELTA": ChangeType_ResubDelta,
	"HEARTBEAT":   ChangeType_Heartbeat,
}

// MarshalJSON marshals the enum as a quoted json string
func (ct ChangeType) MarshalJSON() ([]byte, error) {
	elem, ok := changeTypeToString[ct]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ct *ChangeType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := changeTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ct = result
	return nil
}

// SegmentType ENUM

type SegmentType int

const (
	SegmentType_SegStart SegmentType = iota + 1
	SegmentType_Seg
	SegmentType_SegEnd
)

func (st SegmentType) String() string {
	return segmentTypeToString[st]
}

var segmentTypeToString = map[SegmentType]string{
	SegmentType_SegStart: "SEG_START",
	SegmentType_Seg:      "SEG",
	SegmentType_SegEnd:   "SEG_END",
}

var segmentTypeToEnum = map[string]SegmentType{
	"SEG_START": SegmentType_SegStart,
	"SEG":       SegmentType_Seg,
	"SEG_END":   SegmentType_SegEnd,
}

// MarshalJSON marshals the enum as a quoted json string
func (st SegmentType) MarshalJSON() ([]byte, error) {
	elem, ok := segmentTypeToString[st]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (st *SegmentType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := segmentTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*st = result
	return nil
}

// RaceStatus ENUM

type RaceStatus int

const (
	RaceStatus_Inactive RaceStatus = iota + 1
	RaceStatus_Open
	RaceStatus_Suspended
	RaceStatus_Closed
)

func (rs RaceStatus) String() string {
	return raceStatusToString[rs]
}

var raceStatusToString = map[RaceStatus]string{
	RaceStatus_Inactive:  "INACTIVE",
	RaceStatus_Open:      "OPEN",
	RaceStatus_Suspended: "SUSPENDED",
	RaceStatus_Closed:    "CLOSED",
}

var raceStatusToEnum = map[string]RaceStatus{
	"INACTIVE":  RaceStatus_Inactive,
	"OPEN":      RaceStatus_Open,
	"SUSPENDED": RaceStatus_Suspended,
	"CLOSED":    RaceStatus_Closed,
}

// MarshalJSON marshals the enum as a quoted json string
func (rs RaceStatus) MarshalJSON() ([]byte, error) {
	elem, ok := raceStatusToString[rs]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (rs *RaceStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := raceStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*rs = result
	return nil
}

// PriceLadderType ENUM

type PriceLadderType int

const (
	PriceLadderType_Classic PriceLadderType = iota + 1
	PriceLadderType_Finest
	PriceLadderType_LineRange
)

func (plt PriceLadderType) String() string {
	return priceLadderTypeToString[plt]
}

var priceLadderTypeToString = map[PriceLadderType]string{
	PriceLadderType_Classic:   "CLASSIC",
	PriceLadderType_Finest:    "FINEST",
	PriceLadderType_LineRange: "LINE_RANGE",
}

var priceLadderTypeToEnum = map[string]PriceLadderType{
	"CLASSIC":    PriceLadderType_Classic,
	"FINEST":     PriceLadderType_Finest,
	"LINE_RANGE": PriceLadderType_LineRange,
}

// MarshalJSON marshals the enum as a quoted json string
func (plt PriceLadderType) MarshalJSON() ([]byte, error) {
	elem, ok := priceLadderTypeToString[plt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (plt *PriceLadderType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := priceLadderTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*plt = result
	return nil
}

// RunnerStatus ENUM

type RunnerStatus int

const (
	RunnerStatus_Active RunnerStatus = iota + 1
	RunnerStatus_Winner
	RunnerStatus_Loser
	RunnerStatus_Removed
	RunnerStatus_RemovedVacant
	RunnerStatus_Hidden
	RunnerStatus_Placed
)

func (rs RunnerStatus) String() string {
	return runnerStatusToString[rs]
}

var runnerStatusToString = map[RunnerStatus]string{
	RunnerStatus_Active:        "ACTIVE",
	RunnerStatus_Winner:        "WINNER",
	RunnerStatus_Loser:         "LOSER",
	RunnerStatus_Removed:       "REMOVED",
	RunnerStatus_RemovedVacant: "REMOVED_VACANT",
	RunnerStatus_Hidden:        "HIDDEN",
	RunnerStatus_Placed:        "PLACED",
}

var runnerStatusToEnum = map[string]RunnerStatus{
	"ACTIVE":         RunnerStatus_Active,
	"WINNER":         RunnerStatus_Winner,
	"LOSER":          RunnerStatus_Loser,
	"REMOVED":        RunnerStatus_Removed,
	"REMOVED_VACANT": RunnerStatus_RemovedVacant,
	"HIDDEN":         RunnerStatus_Hidden,
	"PLACED":         RunnerStatus_Placed,
}

// MarshalJSON marshals the enum as a quoted json string
func (rs RunnerStatus) MarshalJSON() ([]byte, error) {
	elem, ok := runnerStatusToString[rs]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (rs *RunnerStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := runnerStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*rs = result
	return nil
}

// OrderSide ENUM

type OrderSide int

const (
	OrderSide_Back OrderSide = iota + 1
	OrderSide_Lay
)

func (os OrderSide) String() string {
	return orderSideToString[os]
}

var orderSideToString = map[OrderSide]string{
	OrderSide_Back: "B",
	OrderSide_Lay:  "L",
}

var orderSideToEnum = map[string]OrderSide{
	"B": OrderSide_Back,
	"L": OrderSide_Lay,
}

// MarshalJSON marshals the enum as a quoted json string
func (os OrderSide) MarshalJSON() ([]byte, error) {
	elem, ok := orderSideToString[os]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (os *OrderSide) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := orderSideToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*os = result
	return nil
}

// PersistenceType ENUM

type PersistenceType int

const (
	PersistenceType_Lapse PersistenceType = iota + 1
	PersistenceType_Persist
	PersistenceType_MarketOnClose
)

func (pt PersistenceType) String() string {
	return persistenceTypeToString[pt]
}

var persistenceTypeToString = map[PersistenceType]string{
	PersistenceType_Lapse:         "L",
	PersistenceType_Persist:       "P",
	PersistenceType_MarketOnClose: "MOC",
}

var persistenceTypeToEnum = map[string]PersistenceType{
	"L":   PersistenceType_Lapse,
	"P":   PersistenceType_Persist,
	"MOC": PersistenceType_MarketOnClose,
}

// MarshalJSON marshals the enum as a quoted json string
func (pt PersistenceType) MarshalJSON() ([]byte, error) {
	elem, ok := persistenceTypeToString[pt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (pt *PersistenceType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := persistenceTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*pt = result
	return nil
}

// OrderType ENUM

type OrderType int

const (
	OrderType_Limit OrderType = iota + 1
	OrderType_MarketOnClose
	OrderType_LimitOnClose
)

func (ot OrderType) String() string {
	return orderTypeToString[ot]
}

var orderTypeToString = map[OrderType]string{
	OrderType_Limit:         "L",
	OrderType_MarketOnClose: "LOC",
	OrderType_LimitOnClose:  "MOC",
}

var orderTypeToEnum = map[string]OrderType{
	"L":   OrderType_Limit,
	"LOC": OrderType_MarketOnClose,
	"MOC": OrderType_LimitOnClose,
}

// MarshalJSON marshals the enum as a quoted json string
func (ot OrderType) MarshalJSON() ([]byte, error) {
	elem, ok := orderTypeToString[ot]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ot *OrderType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := orderTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ot = result
	return nil
}

// OrderStatus ENUM

type OrderStatus int

const (
	OrderStatus_Executable OrderStatus = iota + 1
	OrderStatus_ExecutableComplete
)

func (os OrderStatus) String() string {
	return orderStatusToString[os]
}

var orderStatusToString = map[OrderStatus]string{
	OrderStatus_Executable:         "E",
	OrderStatus_ExecutableComplete: "EC",
}

var orderStatusToEnum = map[string]OrderStatus{
	"E":  OrderStatus_Executable,
	"EC": OrderStatus_ExecutableComplete,
}

// MarshalJSON marshals the enum as a quoted json string
func (os OrderStatus) MarshalJSON() ([]byte, error) {
	elem, ok := orderStatusToString[os]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (os *OrderStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := orderStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*os = result
	return nil
}
