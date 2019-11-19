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
	return ErrorCodeToString[ec]
}

var ErrorCodeToString = map[ErrorCode]string{
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

var ErrorCodeToEnum = map[string]ErrorCode{
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
	elem, ok := ErrorCodeToString[ec]
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

	result, ok := ErrorCodeToEnum[j]
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
	return StatusCodeToString[sc]
}

var StatusCodeToString = map[StatusCode]string{
	StatusCode_Success: "SUCCESS",
	StatusCode_Failure: "FAILURE",
}

var StatusCodeToEnum = map[string]StatusCode{
	"SUCCESS": StatusCode_Success,
	"FAILURE": StatusCode_Failure,
}

// MarshalJSON marshals the enum as a quoted json string
func (sc StatusCode) MarshalJSON() ([]byte, error) {
	elem, ok := StatusCodeToString[sc]
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

	result, ok := StatusCodeToEnum[j]
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
	return PriceDataToString[pd]
}

var PriceDataToString = map[PriceData]string{
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

var PriceDataToEnum = map[string]PriceData{
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
	elem, ok := PriceDataToString[pd]
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

	result, ok := PriceDataToEnum[j]
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
	return BettingTypeToString[bt]
}

var BettingTypeToString = map[BettingType]string{
	BettingType_Odds:                    "ODDS",
	BettingType_Line:                    "LINE",
	BettingType_Range:                   "RANGE",
	BettingType_AsianHandicapDoubleLine: "ASIAN_HANDICAP_DOUBLE_LINE",
	BettingType_AsianHandicapSingleLine: "ASIAN_HANDICAP_SINGLE_LINE",
}

var BettingTypeToEnum = map[string]BettingType{
	"ODDS":                       BettingType_Odds,
	"LINE":                       BettingType_Line,
	"RANGE":                      BettingType_Range,
	"ASIAN_HANDICAP_DOUBLE_LINE": BettingType_AsianHandicapDoubleLine,
	"ASIAN_HANDICAP_SINGLE_LINE": BettingType_AsianHandicapSingleLine,
}

// MarshalJSON marshals the enum as a quoted json string
func (bt BettingType) MarshalJSON() ([]byte, error) {
	elem, ok := BettingTypeToString[bt]
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

	result, ok := BettingTypeToEnum[j]
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
	return ChangeTypeToString[ct]
}

var ChangeTypeToString = map[ChangeType]string{
	ChangeType_SubImage:   "SUB_IMAGE",
	ChangeType_ResubDelta: "RESUB_DELTA",
	ChangeType_Heartbeat:  "HEARTBEAT",
}

var ChangeTypeToEnum = map[string]ChangeType{
	"SUB_IMAGE":   ChangeType_SubImage,
	"RESUB_DELTA": ChangeType_ResubDelta,
	"HEARTBEAT":   ChangeType_Heartbeat,
}

// MarshalJSON marshals the enum as a quoted json string
func (ct ChangeType) MarshalJSON() ([]byte, error) {
	elem, ok := ChangeTypeToString[ct]
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

	result, ok := ChangeTypeToEnum[j]
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
	return SegmentTypeToString[st]
}

var SegmentTypeToString = map[SegmentType]string{
	SegmentType_SegStart: "SEG_START",
	SegmentType_Seg:      "SEG",
	SegmentType_SegEnd:   "SEG_END",
}

var SegmentTypeToEnum = map[string]SegmentType{
	"SEG_START": SegmentType_SegStart,
	"SEG":       SegmentType_Seg,
	"SEG_END":   SegmentType_SegEnd,
}

// MarshalJSON marshals the enum as a quoted json string
func (st SegmentType) MarshalJSON() ([]byte, error) {
	elem, ok := SegmentTypeToString[st]
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

	result, ok := SegmentTypeToEnum[j]
	if !ok {
		return errors.New("couldn't find matching enum value")
	}

	*st = result
	return nil
}
