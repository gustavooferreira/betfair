package betting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// MarketProjection ENUM

type MarketProjection int

const (
	MarketProjection_Competition MarketProjection = iota + 1
	MarketProjection_Event
	MarketProjection_EventType
	MarketProjection_MarketStartTime
	MarketProjection_MarketDescription
	MarketProjection_RunnerDescription
	MarketProjection_RunnerMetadata
)

func (mp MarketProjection) String() string {
	return marketProjectionToString[mp]
}

var marketProjectionToString = map[MarketProjection]string{
	MarketProjection_Competition:       "COMPETITION",
	MarketProjection_Event:             "EVENT",
	MarketProjection_EventType:         "EVENT_TYPE",
	MarketProjection_MarketStartTime:   "MARKET_START_TIME",
	MarketProjection_MarketDescription: "MARKET_DESCRIPTION",
	MarketProjection_RunnerDescription: "RUNNER_DESCRIPTION",
	MarketProjection_RunnerMetadata:    "RUNNER_METADATA",
}

var marketProjectionToEnum = map[string]MarketProjection{
	"COMPETITION":        MarketProjection_Competition,
	"EVENT":              MarketProjection_Event,
	"EVENT_TYPE":         MarketProjection_EventType,
	"MARKET_START_TIME":  MarketProjection_MarketStartTime,
	"MARKET_DESCRIPTION": MarketProjection_MarketDescription,
	"RUNNER_DESCRIPTION": MarketProjection_RunnerDescription,
	"RUNNER_METADATA":    MarketProjection_RunnerMetadata,
}

// MarshalJSON marshals the enum as a quoted json string
func (mp MarketProjection) MarshalJSON() ([]byte, error) {
	elem, ok := marketProjectionToString[mp]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (mp *MarketProjection) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := marketProjectionToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*mp = result
	return nil
}

// PriceData ENUM

type PriceData int

const (
	PriceData_SpAvailable PriceData = iota + 1
	PriceData_SpTraded
	PriceData_ExBestOffers
	PriceData_ExAllOffers
	PriceData_ExTraded
)

func (pd PriceData) String() string {
	return priceDataToString[pd]
}

var priceDataToString = map[PriceData]string{
	PriceData_SpAvailable:  "SP_AVAILABLE",
	PriceData_SpTraded:     "SP_TRADED",
	PriceData_ExBestOffers: "EX_BEST_OFFERS",
	PriceData_ExAllOffers:  "EX_ALL_OFFERS",
	PriceData_ExTraded:     "EX_TRADED",
}

var priceDataToEnum = map[string]PriceData{
	"SP_AVAILABLE":   PriceData_SpAvailable,
	"SP_TRADED":      PriceData_SpTraded,
	"EX_BEST_OFFERS": PriceData_ExBestOffers,
	"EX_ALL_OFFERS":  PriceData_ExAllOffers,
	"EX_TRADED":      PriceData_ExTraded,
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

// MatchProjection ENUM

type MatchProjection int

const (
	MatchProjection_NoRollup MatchProjection = iota + 1
	MatchProjection_RolledUpByPrice
	MatchProjection_RolledUpByAvgPrice
)

func (mp MatchProjection) String() string {
	return matchProjectionToString[mp]
}

var matchProjectionToString = map[MatchProjection]string{
	MatchProjection_NoRollup:           "NO_ROLLUP",
	MatchProjection_RolledUpByPrice:    "ROLLED_UP_BY_PRICE",
	MatchProjection_RolledUpByAvgPrice: "ROLLED_UP_BY_AVG_PRICE",
}

var matchProjectionToEnum = map[string]MatchProjection{
	"NO_ROLLUP":              MatchProjection_NoRollup,
	"ROLLED_UP_BY_PRICE":     MatchProjection_RolledUpByPrice,
	"ROLLED_UP_BY_AVG_PRICE": MatchProjection_RolledUpByAvgPrice,
}

// MarshalJSON marshals the enum as a quoted json string
func (mp MatchProjection) MarshalJSON() ([]byte, error) {
	elem, ok := matchProjectionToString[mp]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (mp *MatchProjection) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := matchProjectionToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*mp = result
	return nil
}

// OrderProjection ENUM

type OrderProjection int

const (
	OrderProjection_All OrderProjection = iota + 1
	OrderProjection_Executable
	OrderProjection_ExecutionComplete
)

func (op OrderProjection) String() string {
	return orderProjectionToString[op]
}

var orderProjectionToString = map[OrderProjection]string{
	OrderProjection_All:               "ALL",
	OrderProjection_Executable:        "EXECUTABLE",
	OrderProjection_ExecutionComplete: "EXECUTION_COMPLETE",
}

var orderProjectionToEnum = map[string]OrderProjection{
	"ALL":                OrderProjection_All,
	"EXECUTABLE":         OrderProjection_Executable,
	"EXECUTION_COMPLETE": OrderProjection_ExecutionComplete,
}

// MarshalJSON marshals the enum as a quoted json string
func (op OrderProjection) MarshalJSON() ([]byte, error) {
	elem, ok := orderProjectionToString[op]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (op *OrderProjection) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := orderProjectionToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*op = result
	return nil
}

// MarketStatus ENUM

type MarketStatus int

const (
	MarketStatus_Inactive MarketStatus = iota + 1
	MarketStatus_Open
	MarketStatus_Suspended
	MarketStatus_Closed
)

func (ms MarketStatus) String() string {
	return marketStatusToString[ms]
}

var marketStatusToString = map[MarketStatus]string{
	MarketStatus_Inactive:  "INACTIVE",
	MarketStatus_Open:      "OPEN",
	MarketStatus_Suspended: "SUSPENDED",
	MarketStatus_Closed:    "CLOSED",
}

var marketStatusToEnum = map[string]MarketStatus{
	"INACTIVE":  MarketStatus_Inactive,
	"OPEN":      MarketStatus_Open,
	"SUSPENDED": MarketStatus_Suspended,
	"CLOSED":    MarketStatus_Closed,
}

// MarshalJSON marshals the enum as a quoted json string
func (ms MarketStatus) MarshalJSON() ([]byte, error) {
	elem, ok := marketStatusToString[ms]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ms *MarketStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := marketStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ms = result
	return nil
}

// RunnerStatus ENUM

type RunnerStatus int

const (
	RunnerStatus_Active RunnerStatus = iota + 1
	RunnerStatus_Winner
	RunnerStatus_Loser
	RunnerStatus_Placed
	RunnerStatus_RemovedVacant
	RunnerStatus_Removed
	RunnerStatus_Hidden
)

func (rs RunnerStatus) String() string {
	return runnerStatusToString[rs]
}

var runnerStatusToString = map[RunnerStatus]string{
	RunnerStatus_Active:        "ACTIVE",
	RunnerStatus_Winner:        "WINNER",
	RunnerStatus_Loser:         "LOSER",
	RunnerStatus_Placed:        "PLACED",
	RunnerStatus_RemovedVacant: "REMOVED_VACANT",
	RunnerStatus_Removed:       "REMOVED",
	RunnerStatus_Hidden:        "HIDDEN",
}

var runnerStatusToEnum = map[string]RunnerStatus{
	"ACTIVE":         RunnerStatus_Active,
	"WINNER":         RunnerStatus_Winner,
	"LOSER":          RunnerStatus_Loser,
	"PLACED":         RunnerStatus_Placed,
	"REMOVED_VACANT": RunnerStatus_RemovedVacant,
	"REMOVED":        RunnerStatus_Removed,
	"HIDDEN":         RunnerStatus_Hidden,
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

// TimeGranularity ENUM

type TimeGranularity int

const (
	TimeGranularity_Days TimeGranularity = iota + 1
	TimeGranularity_Hours
	TimeGranularity_Minutes
)

func (tg TimeGranularity) String() string {
	return timeGranularityToString[tg]
}

var timeGranularityToString = map[TimeGranularity]string{
	TimeGranularity_Days:    "DAYS",
	TimeGranularity_Hours:   "HOURS",
	TimeGranularity_Minutes: "MINUTES",
}

var timeGranularityToEnum = map[string]TimeGranularity{
	"DAYS":    TimeGranularity_Days,
	"HOURS":   TimeGranularity_Hours,
	"MINUTES": TimeGranularity_Minutes,
}

// MarshalJSON marshals the enum as a quoted json string
func (tg TimeGranularity) MarshalJSON() ([]byte, error) {
	elem, ok := timeGranularityToString[tg]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (tg *TimeGranularity) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := timeGranularityToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*tg = result
	return nil
}

// Side ENUM

type Side int

const (
	Side_Back Side = iota + 1
	Side_Lay
)

func (s Side) String() string {
	return sideToString[s]
}

var sideToString = map[Side]string{
	Side_Back: "BACK",
	Side_Lay:  "LAY",
}

var sideToEnum = map[string]Side{
	"BACK": Side_Back,
	"LAY":  Side_Lay,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Side) MarshalJSON() ([]byte, error) {
	elem, ok := sideToString[s]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Side) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := sideToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*s = result
	return nil
}

// OrderStatus ENUM

type OrderStatus int

const (
	OrderStatus_Pending OrderStatus = iota + 1
	OrderStatus_ExecutionComplete
	OrderStatus_Executable
	OrderStatus_Expired
)

func (os OrderStatus) String() string {
	return orderStatusToString[os]
}

var orderStatusToString = map[OrderStatus]string{
	OrderStatus_Pending:           "PENDING",
	OrderStatus_ExecutionComplete: "EXECUTION_COMPLETE",
	OrderStatus_Executable:        "EXECUTABLE",
	OrderStatus_Expired:           "EXPIRED",
}

var orderStatusToEnum = map[string]OrderStatus{
	"PENDING":            OrderStatus_Pending,
	"EXECUTION_COMPLETE": OrderStatus_ExecutionComplete,
	"EXECUTABLE":         OrderStatus_Executable,
	"EXPIRED":            OrderStatus_Expired,
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

// OrderBy ENUM

type OrderBy int

const (
	OrderBy_ByBet OrderBy = iota + 1
	OrderBy_ByMarket
	OrderBy_ByMatchTime
	OrderBy_ByPlaceTime
	OrderBy_BySettledTime
	OrderBy_ByVoidTime
)

func (ob OrderBy) String() string {
	return orderByToString[ob]
}

var orderByToString = map[OrderBy]string{
	OrderBy_ByBet:         "BY_BET",
	OrderBy_ByMarket:      "BY_MARKET",
	OrderBy_ByMatchTime:   "BY_MATCH_TIME",
	OrderBy_ByPlaceTime:   "BY_PLACE_TIME",
	OrderBy_BySettledTime: "BY_SETTLED_TIME",
	OrderBy_ByVoidTime:    "BY_VOID_TIME",
}

var orderByToEnum = map[string]OrderBy{
	"BY_BET":          OrderBy_ByBet,
	"BY_MARKET":       OrderBy_ByMarket,
	"BY_MATCH_TIME":   OrderBy_ByMatchTime,
	"BY_PLACE_TIME":   OrderBy_ByPlaceTime,
	"BY_SETTLED_TIME": OrderBy_BySettledTime,
	"BY_VOID_TIME":    OrderBy_ByVoidTime,
}

// MarshalJSON marshals the enum as a quoted json string
func (ob OrderBy) MarshalJSON() ([]byte, error) {
	elem, ok := orderByToString[ob]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ob *OrderBy) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := orderByToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ob = result
	return nil
}

// SortDir ENUM

type SortDir int

const (
	SortDir_EarliestToLatest SortDir = iota + 1
	SortDir_LatestToEarliest
)

func (sd SortDir) String() string {
	return sortDirToString[sd]
}

var sortDirToString = map[SortDir]string{
	SortDir_EarliestToLatest: "EARLIEST_TO_LATEST",
	SortDir_LatestToEarliest: "LATEST_TO_EARLIEST",
}

var sortDirToEnum = map[string]SortDir{
	"EARLIEST_TO_LATEST": SortDir_EarliestToLatest,
	"LATEST_TO_EARLIEST": SortDir_LatestToEarliest,
}

// MarshalJSON marshals the enum as a quoted json string
func (sd SortDir) MarshalJSON() ([]byte, error) {
	elem, ok := sortDirToString[sd]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (sd *SortDir) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := sortDirToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*sd = result
	return nil
}

// OrderType ENUM

type OrderType int

const (
	OrderType_Limit OrderType = iota + 1
	OrderType_LimitOnClose
	OrderType_MarketOnClose
)

func (ot OrderType) String() string {
	return orderTypeToString[ot]
}

var orderTypeToString = map[OrderType]string{
	OrderType_Limit:         "LIMIT",
	OrderType_LimitOnClose:  "LIMIT_ON_CLOSE",
	OrderType_MarketOnClose: "MARKET_ON_CLOSE",
}

var orderTypeToEnum = map[string]OrderType{
	"LIMIT":           OrderType_Limit,
	"LIMIT_ON_CLOSE":  OrderType_LimitOnClose,
	"MARKET_ON_CLOSE": OrderType_MarketOnClose,
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

// MarketSort ENUM

type MarketSort int

const (
	MarketSort_MinimumTraded MarketSort = iota + 1
	MarketSort_MaximumTraded
	MarketSort_MinimumAvailable
	MarketSort_MaximumAvailable
	MarketSort_FirstToStart
	MarketSort_LastToStart
)

func (ms MarketSort) String() string {
	return marketSortToString[ms]
}

var marketSortToString = map[MarketSort]string{
	MarketSort_MinimumTraded:    "MINIMUM_TRADED",
	MarketSort_MaximumTraded:    "MAXIMUM_TRADED",
	MarketSort_MinimumAvailable: "MINIMUM_AVAILABLE",
	MarketSort_MaximumAvailable: "MAXIMUM_AVAILABLE",
	MarketSort_FirstToStart:     "FIRST_TO_START",
	MarketSort_LastToStart:      "LAST_TO_START",
}

var marketSortToEnum = map[string]MarketSort{
	"MINIMUM_TRADED":    MarketSort_MinimumTraded,
	"MAXIMUM_TRADED":    MarketSort_MaximumTraded,
	"MINIMUM_AVAILABLE": MarketSort_MinimumAvailable,
	"MAXIMUM_AVAILABLE": MarketSort_MaximumAvailable,
	"FIRST_TO_START":    MarketSort_FirstToStart,
	"LAST_TO_START":     MarketSort_LastToStart,
}

// MarshalJSON marshals the enum as a quoted json string
func (ms MarketSort) MarshalJSON() ([]byte, error) {
	elem, ok := marketSortToString[ms]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ms *MarketSort) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := marketSortToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ms = result
	return nil
}

// MarketBettingType ENUM

type MarketBettingType int

const (
	MarketBettingType_Odds MarketBettingType = iota + 1
	MarketBettingType_Line
	MarketBettingType_Range
	MarketBettingType_AsianHandicapDoubleLine
	MarketBettingType_AsianHandicapSingleLine
	MarketBettingType_FixedOdds
)

func (mbt MarketBettingType) String() string {
	return marketBettingTypeToString[mbt]
}

var marketBettingTypeToString = map[MarketBettingType]string{
	MarketBettingType_Odds:                    "ODDS",
	MarketBettingType_Line:                    "LINE",
	MarketBettingType_Range:                   "RANGE",
	MarketBettingType_AsianHandicapDoubleLine: "ASIAN_HANDICAP_DOUBLE_LINE",
	MarketBettingType_AsianHandicapSingleLine: "ASIAN_HANDICAP_SINGLE_LINE",
	MarketBettingType_FixedOdds:               "FIXED_ODDS",
}

var marketBettingTypeToEnum = map[string]MarketBettingType{
	"ODDS":                       MarketBettingType_Odds,
	"LINE":                       MarketBettingType_Line,
	"RANGE":                      MarketBettingType_Range,
	"ASIAN_HANDICAP_DOUBLE_LINE": MarketBettingType_AsianHandicapDoubleLine,
	"ASIAN_HANDICAP_SINGLE_LINE": MarketBettingType_AsianHandicapSingleLine,
	"FIXED_ODDS":                 MarketBettingType_FixedOdds,
}

// MarshalJSON marshals the enum as a quoted json string
func (mbt MarketBettingType) MarshalJSON() ([]byte, error) {
	elem, ok := marketBettingTypeToString[mbt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (mbt *MarketBettingType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := marketBettingTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*mbt = result
	return nil
}

// ExecutionReportStatus ENUM

type ExecutionReportStatus int

const (
	ExecutionReportStatus_Success ExecutionReportStatus = iota + 1
	ExecutionReportStatus_Failure
	ExecutionReportStatus_ProcessedWithErrors
	ExecutionReportStatus_Timeout
)

func (ers ExecutionReportStatus) String() string {
	return executionReportStatusToString[ers]
}

var executionReportStatusToString = map[ExecutionReportStatus]string{
	ExecutionReportStatus_Success:             "SUCCESS",
	ExecutionReportStatus_Failure:             "FAILURE",
	ExecutionReportStatus_ProcessedWithErrors: "PROCESSED_WITH_ERRORS",
	ExecutionReportStatus_Timeout:             "TIMEOUT",
}

var executionReportStatusToEnum = map[string]ExecutionReportStatus{
	"SUCCESS":               ExecutionReportStatus_Success,
	"FAILURE":               ExecutionReportStatus_Failure,
	"PROCESSED_WITH_ERRORS": ExecutionReportStatus_ProcessedWithErrors,
	"TIMEOUT":               ExecutionReportStatus_Timeout,
}

// MarshalJSON marshals the enum as a quoted json string
func (ers ExecutionReportStatus) MarshalJSON() ([]byte, error) {
	elem, ok := executionReportStatusToString[ers]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ers *ExecutionReportStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := executionReportStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*ers = result
	return nil
}

// ExecutionReportErrorCode ENUM

type ExecutionReportErrorCode int

const (
	ExecutionReportErrorCode_ErrorInMatcher ExecutionReportErrorCode = iota + 1
	ExecutionReportErrorCode_ProcessedWithErrors
	ExecutionReportErrorCode_BetActionError
	ExecutionReportErrorCode_InvalidAccountState
	ExecutionReportErrorCode_InvalidWalletStatus
	ExecutionReportErrorCode_InsufficientFunds
	ExecutionReportErrorCode_LossLimitExceeded
	ExecutionReportErrorCode_MarketSuspended
	ExecutionReportErrorCode_MarketNotOpenForBetting
	ExecutionReportErrorCode_DuplicateTransaction
	ExecutionReportErrorCode_InvalidOrder
	ExecutionReportErrorCode_InvalidMarketId
	ExecutionReportErrorCode_PermissionDenied
	ExecutionReportErrorCode_DuplicateBetids
	ExecutionReportErrorCode_NoActionRequired
	ExecutionReportErrorCode_ServiceUnavailable
	ExecutionReportErrorCode_RejectedByRegulator
	ExecutionReportErrorCode_NoChasing
	ExecutionReportErrorCode_RegulatorIsNotAvailable
	ExecutionReportErrorCode_TooManyInstructions
	ExecutionReportErrorCode_InvalidMarketVersion
)

func (erec ExecutionReportErrorCode) String() string {
	return executionReportErrorCodeToString[erec]
}

var executionReportErrorCodeToString = map[ExecutionReportErrorCode]string{
	ExecutionReportErrorCode_ErrorInMatcher:          "ERROR_IN_MATCHER",
	ExecutionReportErrorCode_ProcessedWithErrors:     "PROCESSED_WITH_ERRORS",
	ExecutionReportErrorCode_BetActionError:          "BET_ACTION_ERROR",
	ExecutionReportErrorCode_InvalidAccountState:     "INVALID_ACCOUNT_STATE",
	ExecutionReportErrorCode_InvalidWalletStatus:     "INVALID_WALLET_STATUS",
	ExecutionReportErrorCode_InsufficientFunds:       "INSUFFICIENT_FUNDS",
	ExecutionReportErrorCode_LossLimitExceeded:       "LOSS_LIMIT_EXCEEDED",
	ExecutionReportErrorCode_MarketSuspended:         "MARKET_SUSPENDED",
	ExecutionReportErrorCode_MarketNotOpenForBetting: "MARKET_NOT_OPEN_FOR_BETTING",
	ExecutionReportErrorCode_DuplicateTransaction:    "DUPLICATE_TRANSACTION",
	ExecutionReportErrorCode_InvalidOrder:            "INVALID_ORDER",
	ExecutionReportErrorCode_InvalidMarketId:         "INVALID_MARKET_ID",
	ExecutionReportErrorCode_PermissionDenied:        "PERMISSION_DENIED",
	ExecutionReportErrorCode_DuplicateBetids:         "DUPLICATE_BETIDS",
	ExecutionReportErrorCode_NoActionRequired:        "NO_ACTION_REQUIRED",
	ExecutionReportErrorCode_ServiceUnavailable:      "SERVICE_UNAVAILABLE",
	ExecutionReportErrorCode_RejectedByRegulator:     "REJECTED_BY_REGULATOR",
	ExecutionReportErrorCode_NoChasing:               "NO_CHASING",
	ExecutionReportErrorCode_RegulatorIsNotAvailable: "REGULATOR_IS_NOT_AVAILABLE",
	ExecutionReportErrorCode_TooManyInstructions:     "TOO_MANY_INSTRUCTIONS",
	ExecutionReportErrorCode_InvalidMarketVersion:    "INVALID_MARKET_VERSION",
}

var executionReportErrorCodeToEnum = map[string]ExecutionReportErrorCode{
	"ERROR_IN_MATCHER":            ExecutionReportErrorCode_ErrorInMatcher,
	"PROCESSED_WITH_ERRORS":       ExecutionReportErrorCode_ProcessedWithErrors,
	"BET_ACTION_ERROR":            ExecutionReportErrorCode_BetActionError,
	"INVALID_ACCOUNT_STATE":       ExecutionReportErrorCode_InvalidAccountState,
	"INVALID_WALLET_STATUS":       ExecutionReportErrorCode_InvalidWalletStatus,
	"INSUFFICIENT_FUNDS":          ExecutionReportErrorCode_InsufficientFunds,
	"LOSS_LIMIT_EXCEEDED":         ExecutionReportErrorCode_LossLimitExceeded,
	"MARKET_SUSPENDED":            ExecutionReportErrorCode_MarketSuspended,
	"MARKET_NOT_OPEN_FOR_BETTING": ExecutionReportErrorCode_MarketNotOpenForBetting,
	"DUPLICATE_TRANSACTION":       ExecutionReportErrorCode_DuplicateTransaction,
	"INVALID_ORDER":               ExecutionReportErrorCode_InvalidOrder,
	"INVALID_MARKET_ID":           ExecutionReportErrorCode_InvalidMarketId,
	"PERMISSION_DENIED":           ExecutionReportErrorCode_PermissionDenied,
	"DUPLICATE_BETIDS":            ExecutionReportErrorCode_DuplicateBetids,
	"NO_ACTION_REQUIRED":          ExecutionReportErrorCode_NoActionRequired,
	"SERVICE_UNAVAILABLE":         ExecutionReportErrorCode_ServiceUnavailable,
	"REJECTED_BY_REGULATOR":       ExecutionReportErrorCode_RejectedByRegulator,
	"NO_CHASING":                  ExecutionReportErrorCode_NoChasing,
	"REGULATOR_IS_NOT_AVAILABLE":  ExecutionReportErrorCode_RegulatorIsNotAvailable,
	"TOO_MANY_INSTRUCTIONS":       ExecutionReportErrorCode_TooManyInstructions,
	"INVALID_MARKET_VERSION":      ExecutionReportErrorCode_InvalidMarketVersion,
}

// MarshalJSON marshals the enum as a quoted json string
func (erec ExecutionReportErrorCode) MarshalJSON() ([]byte, error) {
	elem, ok := executionReportErrorCodeToString[erec]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (erec *ExecutionReportErrorCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := executionReportErrorCodeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*erec = result
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
	PersistenceType_Lapse:         "LAPSE",
	PersistenceType_Persist:       "PERSIST",
	PersistenceType_MarketOnClose: "MARKET_ON_CLOSE",
}

var persistenceTypeToEnum = map[string]PersistenceType{
	"LAPSE":           PersistenceType_Lapse,
	"PERSIST":         PersistenceType_Persist,
	"MARKET_ON_CLOSE": PersistenceType_MarketOnClose,
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

// InstructionReportStatus ENUM

type InstructionReportStatus int

const (
	InstructionReportStatus_Success InstructionReportStatus = iota + 1
	InstructionReportStatus_Failure
	InstructionReportStatus_Timeout
)

func (irs InstructionReportStatus) String() string {
	return instructionReportStatusToString[irs]
}

var instructionReportStatusToString = map[InstructionReportStatus]string{
	InstructionReportStatus_Success: "SUCCESS",
	InstructionReportStatus_Failure: "FAILURE",
	InstructionReportStatus_Timeout: "TIMEOUT",
}

var instructionReportStatusToEnum = map[string]InstructionReportStatus{
	"SUCCESS": InstructionReportStatus_Success,
	"FAILURE": InstructionReportStatus_Failure,
	"TIMEOUT": InstructionReportStatus_Timeout,
}

// MarshalJSON marshals the enum as a quoted json string
func (irs InstructionReportStatus) MarshalJSON() ([]byte, error) {
	elem, ok := instructionReportStatusToString[irs]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (irs *InstructionReportStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := instructionReportStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*irs = result
	return nil
}

// InstructionReportErrorCode ENUM

type InstructionReportErrorCode int

const (
	InstructionReportErrorCode_InvalidBetSize InstructionReportErrorCode = iota + 1
	InstructionReportErrorCode_InvalidRunner
	InstructionReportErrorCode_BetTakenOrLapsed
	InstructionReportErrorCode_BetInProgress
	InstructionReportErrorCode_RunnerRemoved
	InstructionReportErrorCode_MarketNotOpenForBetting
	InstructionReportErrorCode_LossLimitExceeded
	InstructionReportErrorCode_MarketNotOpenForBspBetting
	InstructionReportErrorCode_InvalidPriceEdit
	InstructionReportErrorCode_InvalidOdds
	InstructionReportErrorCode_InsufficientFunds
	InstructionReportErrorCode_InvalidPersistenceType
	InstructionReportErrorCode_ErrorInMatcher
	InstructionReportErrorCode_InvalidBackLayCombination
	InstructionReportErrorCode_ErrorInOrder
	InstructionReportErrorCode_InvalidBidType
	InstructionReportErrorCode_InvalidBetId
	InstructionReportErrorCode_CancelledNotPlaced
	InstructionReportErrorCode_RelatedActionFailed
	InstructionReportErrorCode_NoActionRequired
	InstructionReportErrorCode_TimeInForceConflict
	InstructionReportErrorCode_UnexpectedPersistenceType
	InstructionReportErrorCode_InvalidOrderType
	InstructionReportErrorCode_UnexpectedMinFillSize
	InstructionReportErrorCode_InvalidCustomerOrderRef
	InstructionReportErrorCode_InvalidMinFillSize
	InstructionReportErrorCode_BetLapsedPriceImprovementTooLarge
)

func (irec InstructionReportErrorCode) String() string {
	return instructionReportErrorCodeToString[irec]
}

var instructionReportErrorCodeToString = map[InstructionReportErrorCode]string{
	InstructionReportErrorCode_InvalidBetSize:                    "INVALID_BET_SIZE",
	InstructionReportErrorCode_InvalidRunner:                     "INVALID_RUNNER",
	InstructionReportErrorCode_BetTakenOrLapsed:                  "BET_TAKEN_OR_LAPSED",
	InstructionReportErrorCode_BetInProgress:                     "BET_IN_PROGRESS",
	InstructionReportErrorCode_RunnerRemoved:                     "RUNNER_REMOVED",
	InstructionReportErrorCode_MarketNotOpenForBetting:           "MARKET_NOT_OPEN_FOR_BETTING",
	InstructionReportErrorCode_LossLimitExceeded:                 "LOSS_LIMIT_EXCEEDED",
	InstructionReportErrorCode_MarketNotOpenForBspBetting:        "MARKET_NOT_OPEN_FOR_BSP_BETTING",
	InstructionReportErrorCode_InvalidPriceEdit:                  "INVALID_PRICE_EDIT",
	InstructionReportErrorCode_InvalidOdds:                       "INVALID_ODDS",
	InstructionReportErrorCode_InsufficientFunds:                 "INSUFFICIENT_FUNDS",
	InstructionReportErrorCode_InvalidPersistenceType:            "INVALID_PERSISTENCE_TYPE",
	InstructionReportErrorCode_ErrorInMatcher:                    "ERROR_IN_MATCHER",
	InstructionReportErrorCode_InvalidBackLayCombination:         "INVALID_BACK_LAY_COMBINATION",
	InstructionReportErrorCode_ErrorInOrder:                      "ERROR_IN_ORDER",
	InstructionReportErrorCode_InvalidBidType:                    "INVALID_BID_TYPE",
	InstructionReportErrorCode_InvalidBetId:                      "INVALID_BET_ID",
	InstructionReportErrorCode_CancelledNotPlaced:                "CANCELLED_NOT_PLACED",
	InstructionReportErrorCode_RelatedActionFailed:               "RELATED_ACTION_FAILED",
	InstructionReportErrorCode_NoActionRequired:                  "NO_ACTION_REQUIRED",
	InstructionReportErrorCode_TimeInForceConflict:               "TIME_IN_FORCE_CONFLICT",
	InstructionReportErrorCode_UnexpectedPersistenceType:         "UNEXPECTED_PERSISTENCE_TYPE",
	InstructionReportErrorCode_InvalidOrderType:                  "INVALID_ORDER_TYPE",
	InstructionReportErrorCode_UnexpectedMinFillSize:             "UNEXPECTED_MIN_FILL_SIZE",
	InstructionReportErrorCode_InvalidCustomerOrderRef:           "INVALID_CUSTOMER_ORDER_REF",
	InstructionReportErrorCode_InvalidMinFillSize:                "INVALID_MIN_FILL_SIZE",
	InstructionReportErrorCode_BetLapsedPriceImprovementTooLarge: "BET_LAPSED_PRICE_IMPROVEMENT_TOO_LARGE",
}

var instructionReportErrorCodeToEnum = map[string]InstructionReportErrorCode{
	"INVALID_BET_SIZE":                       InstructionReportErrorCode_InvalidBetSize,
	"INVALID_RUNNER":                         InstructionReportErrorCode_InvalidRunner,
	"BET_TAKEN_OR_LAPSED":                    InstructionReportErrorCode_BetTakenOrLapsed,
	"BET_IN_PROGRESS":                        InstructionReportErrorCode_BetInProgress,
	"RUNNER_REMOVED":                         InstructionReportErrorCode_RunnerRemoved,
	"MARKET_NOT_OPEN_FOR_BETTING":            InstructionReportErrorCode_MarketNotOpenForBetting,
	"LOSS_LIMIT_EXCEEDED":                    InstructionReportErrorCode_LossLimitExceeded,
	"MARKET_NOT_OPEN_FOR_BSP_BETTING":        InstructionReportErrorCode_MarketNotOpenForBspBetting,
	"INVALID_PRICE_EDIT":                     InstructionReportErrorCode_InvalidPriceEdit,
	"INVALID_ODDS":                           InstructionReportErrorCode_InvalidOdds,
	"INSUFFICIENT_FUNDS":                     InstructionReportErrorCode_InsufficientFunds,
	"INVALID_PERSISTENCE_TYPE":               InstructionReportErrorCode_InvalidPersistenceType,
	"ERROR_IN_MATCHER":                       InstructionReportErrorCode_ErrorInMatcher,
	"INVALID_BACK_LAY_COMBINATION":           InstructionReportErrorCode_InvalidBackLayCombination,
	"ERROR_IN_ORDER":                         InstructionReportErrorCode_ErrorInOrder,
	"INVALID_BID_TYPE":                       InstructionReportErrorCode_InvalidBidType,
	"INVALID_BET_ID":                         InstructionReportErrorCode_InvalidBetId,
	"CANCELLED_NOT_PLACED":                   InstructionReportErrorCode_CancelledNotPlaced,
	"RELATED_ACTION_FAILED":                  InstructionReportErrorCode_RelatedActionFailed,
	"NO_ACTION_REQUIRED":                     InstructionReportErrorCode_NoActionRequired,
	"TIME_IN_FORCE_CONFLICT":                 InstructionReportErrorCode_TimeInForceConflict,
	"UNEXPECTED_PERSISTENCE_TYPE":            InstructionReportErrorCode_UnexpectedPersistenceType,
	"INVALID_ORDER_TYPE":                     InstructionReportErrorCode_InvalidOrderType,
	"UNEXPECTED_MIN_FILL_SIZE":               InstructionReportErrorCode_UnexpectedMinFillSize,
	"INVALID_CUSTOMER_ORDER_REF":             InstructionReportErrorCode_InvalidCustomerOrderRef,
	"INVALID_MIN_FILL_SIZE":                  InstructionReportErrorCode_InvalidMinFillSize,
	"BET_LAPSED_PRICE_IMPROVEMENT_TOO_LARGE": InstructionReportErrorCode_BetLapsedPriceImprovementTooLarge,
}

// MarshalJSON marshals the enum as a quoted json string
func (irec InstructionReportErrorCode) MarshalJSON() ([]byte, error) {
	elem, ok := instructionReportErrorCodeToString[irec]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (irec *InstructionReportErrorCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := instructionReportErrorCodeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*irec = result
	return nil
}

// RollupModel ENUM

type RollupModel int

const (
	RollupModel_Stake RollupModel = iota + 1
	RollupModel_Payout
	RollupModel_ManagedLiability
	RollupModel_None
)

func (rm RollupModel) String() string {
	return rollupModelToString[rm]
}

var rollupModelToString = map[RollupModel]string{
	RollupModel_Stake:            "STAKE",
	RollupModel_Payout:           "PAYOUT",
	RollupModel_ManagedLiability: "MANAGED_LIABILITY",
	RollupModel_None:             "NONE",
}

var rollupModelToEnum = map[string]RollupModel{
	"STAKE":             RollupModel_Stake,
	"PAYOUT":            RollupModel_Payout,
	"MANAGED_LIABILITY": RollupModel_ManagedLiability,
	"NONE":              RollupModel_None,
}

// MarshalJSON marshals the enum as a quoted json string
func (rm RollupModel) MarshalJSON() ([]byte, error) {
	elem, ok := rollupModelToString[rm]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (rm *RollupModel) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := rollupModelToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*rm = result
	return nil
}

// GroupBy ENUM

type GroupBy int

const (
	GroupBy_EventType GroupBy = iota + 1
	GroupBy_Event
	GroupBy_Market
	GroupBy_Side
	GroupBy_Bet
)

func (gb GroupBy) String() string {
	return groupByToString[gb]
}

var groupByToString = map[GroupBy]string{
	GroupBy_EventType: "EVENT_TYPE",
	GroupBy_Event:     "EVENT",
	GroupBy_Market:    "MARKET",
	GroupBy_Side:      "SIDE",
	GroupBy_Bet:       "BET",
}

var groupByToEnum = map[string]GroupBy{
	"EVENT_TYPE": GroupBy_EventType,
	"EVENT":      GroupBy_Event,
	"MARKET":     GroupBy_Market,
	"SIDE":       GroupBy_Side,
	"BET":        GroupBy_Bet,
}

// MarshalJSON marshals the enum as a quoted json string
func (gb GroupBy) MarshalJSON() ([]byte, error) {
	elem, ok := groupByToString[gb]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (gb *GroupBy) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := groupByToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*gb = result
	return nil
}

// BetStatus ENUM

type BetStatus int

const (
	BetStatus_Settled BetStatus = iota + 1
	BetStatus_Voided
	BetStatus_Lapsed
	BetStatus_Cancelled
)

func (bs BetStatus) String() string {
	return betStatusToString[bs]
}

var betStatusToString = map[BetStatus]string{
	BetStatus_Settled:   "SETTLED",
	BetStatus_Voided:    "VOIDED",
	BetStatus_Lapsed:    "LAPSED",
	BetStatus_Cancelled: "CANCELLED",
}

var betStatusToEnum = map[string]BetStatus{
	"SETTLED":   BetStatus_Settled,
	"VOIDED":    BetStatus_Voided,
	"LAPSED":    BetStatus_Lapsed,
	"CANCELLED": BetStatus_Cancelled,
}

// MarshalJSON marshals the enum as a quoted json string
func (bs BetStatus) MarshalJSON() ([]byte, error) {
	elem, ok := betStatusToString[bs]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (bs *BetStatus) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := betStatusToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*bs = result
	return nil
}

// MarketType ENUM

type MarketType int

const (
	MarketType_A MarketType = iota + 1
	MarketType_L
	MarketType_O
	MarketType_R
	MarketType_NotApplicable
)

func (mt MarketType) String() string {
	return marketTypeToString[mt]
}

var marketTypeToString = map[MarketType]string{
	MarketType_A:             "A",
	MarketType_L:             "L",
	MarketType_O:             "O",
	MarketType_R:             "R",
	MarketType_NotApplicable: "NOT_APPLICABLE",
}

var marketTypeToEnum = map[string]MarketType{
	"A":              MarketType_A,
	"L":              MarketType_L,
	"O":              MarketType_O,
	"R":              MarketType_R,
	"NOT_APPLICABLE": MarketType_NotApplicable,
}

// MarshalJSON marshals the enum as a quoted json string
func (mt MarketType) MarshalJSON() ([]byte, error) {
	elem, ok := marketTypeToString[mt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (mt *MarketType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := marketTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*mt = result
	return nil
}

// TimeInForce ENUM

type TimeInForce int

const (
	TimeInForce_FillOrKill TimeInForce = iota + 1
)

func (tif TimeInForce) String() string {
	return timeInForceToString[tif]
}

var timeInForceToString = map[TimeInForce]string{
	TimeInForce_FillOrKill: "FILL_OR_KILL",
}

var timeInForceToEnum = map[string]TimeInForce{
	"FILL_OR_KILL": TimeInForce_FillOrKill,
}

// MarshalJSON marshals the enum as a quoted json string
func (tif TimeInForce) MarshalJSON() ([]byte, error) {
	elem, ok := timeInForceToString[tif]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (tif *TimeInForce) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := timeInForceToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*tif = result
	return nil
}

// BetTargetType ENUM

type BetTargetType int

const (
	BetTargetType_BackersProfit BetTargetType = iota + 1
	BetTargetType_Payout
)

func (btt BetTargetType) String() string {
	return betTargetTypeToString[btt]
}

var betTargetTypeToString = map[BetTargetType]string{
	BetTargetType_BackersProfit: "BACKERS_PROFIT",
	BetTargetType_Payout:        "PAYOUT",
}

var betTargetTypeToEnum = map[string]BetTargetType{
	"BACKERS_PROFIT": BetTargetType_BackersProfit,
	"PAYOUT":         BetTargetType_Payout,
}

// MarshalJSON marshals the enum as a quoted json string
func (btt BetTargetType) MarshalJSON() ([]byte, error) {
	elem, ok := betTargetTypeToString[btt]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (btt *BetTargetType) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := betTargetTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*btt = result
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

// APINGExceptionCode ENUM

type APINGExceptionCode int

const (
	APINGExceptionCode_TooMuchData APINGExceptionCode = iota + 1
	APINGExceptionCode_InvalidInputData
	APINGExceptionCode_InvalidSessionInformation
	APINGExceptionCode_NoAppKey
	APINGExceptionCode_NoSession
	APINGExceptionCode_UnexpectedError
	APINGExceptionCode_InvalidAppKey
	APINGExceptionCode_TooManyRequests
	APINGExceptionCode_ServiceBusy
	APINGExceptionCode_TimeoutError
	APINGExceptionCode_RequestSizeExceedsLimit
	APINGExceptionCode_AccessDenied
)

func (apingec APINGExceptionCode) String() string {
	return aPINGExceptionCodeToString[apingec]
}

var aPINGExceptionCodeToString = map[APINGExceptionCode]string{
	APINGExceptionCode_TooMuchData:               "TOO_MUCH_DATA",
	APINGExceptionCode_InvalidInputData:          "INVALID_INPUT_DATA",
	APINGExceptionCode_InvalidSessionInformation: "INVALID_SESSION_INFORMATION",
	APINGExceptionCode_NoAppKey:                  "NO_APP_KEY",
	APINGExceptionCode_NoSession:                 "NO_SESSION",
	APINGExceptionCode_UnexpectedError:           "UNEXPECTED_ERROR",
	APINGExceptionCode_InvalidAppKey:             "INVALID_APP_KEY",
	APINGExceptionCode_TooManyRequests:           "TOO_MANY_REQUESTS",
	APINGExceptionCode_ServiceBusy:               "SERVICE_BUSY",
	APINGExceptionCode_TimeoutError:              "TIMEOUT_ERROR",
	APINGExceptionCode_RequestSizeExceedsLimit:   "REQUEST_SIZE_EXCEEDS_LIMIT",
	APINGExceptionCode_AccessDenied:              "ACCESS_DENIED",
}

var aPINGExceptionCodeToEnum = map[string]APINGExceptionCode{
	"TOO_MUCH_DATA":               APINGExceptionCode_TooMuchData,
	"INVALID_INPUT_DATA":          APINGExceptionCode_InvalidInputData,
	"INVALID_SESSION_INFORMATION": APINGExceptionCode_InvalidSessionInformation,
	"NO_APP_KEY":                  APINGExceptionCode_NoAppKey,
	"NO_SESSION":                  APINGExceptionCode_NoSession,
	"UNEXPECTED_ERROR":            APINGExceptionCode_UnexpectedError,
	"INVALID_APP_KEY":             APINGExceptionCode_InvalidAppKey,
	"TOO_MANY_REQUESTS":           APINGExceptionCode_TooManyRequests,
	"SERVICE_BUSY":                APINGExceptionCode_ServiceBusy,
	"TIMEOUT_ERROR":               APINGExceptionCode_TimeoutError,
	"REQUEST_SIZE_EXCEEDS_LIMIT":  APINGExceptionCode_RequestSizeExceedsLimit,
	"ACCESS_DENIED":               APINGExceptionCode_AccessDenied,
}

// MarshalJSON marshals the enum as a quoted json string
func (apingec APINGExceptionCode) MarshalJSON() ([]byte, error) {
	elem, ok := aPINGExceptionCodeToString[apingec]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("couldn't marshal enum")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (apingec *APINGExceptionCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	result, ok := aPINGExceptionCodeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*apingec = result
	return nil
}
