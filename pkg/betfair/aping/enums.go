package aping

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// MarketProjection ENUM

type MarketProjection int

const (
	MarketProjection_Competition MarketProjection = iota
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

	return bytes.NewBufferString("").Bytes(), errors.New("Yolo error")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (mp *MarketProjection) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	// Change this to actually return an error if it can't be found
	*mp = marketProjectionToEnum[j]
	return nil
}

// MarketSort ENUM

type MarketSort int

const (
	MarketSort_FirstToStart MarketSort = iota
)

func (ms MarketSort) String() string {
	return marketSortToString[ms]
}

var marketSortToString = map[MarketSort]string{
	MarketSort_FirstToStart: "FIRST_TO_START",
}

var marketSortToEnum = map[string]MarketSort{
	"FIRST_TO_START": MarketSort_FirstToStart,
}

// MarshalJSON marshals the enum as a quoted json string
func (ms MarketSort) MarshalJSON() ([]byte, error) {
	elem, ok := marketSortToString[ms]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("Yolo error")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ms *MarketSort) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	// Change this to actually return an error if it can't be found
	*ms = marketSortToEnum[j]
	return nil
}

// APINGExceptionCode ENUM

type APINGExceptionCode int

const (
	APINGExceptionCode_InvalidAppKey APINGExceptionCode = iota
)

func (apinge APINGExceptionCode) String() string {
	return apingExceptionToString[apinge]
}

var apingExceptionToString = map[APINGExceptionCode]string{
	APINGExceptionCode_InvalidAppKey: "INVALID_APP_KEY",
}

var apingExceptionToEnum = map[string]APINGExceptionCode{
	"INVALID_APP_KEY": APINGExceptionCode_InvalidAppKey,
}

// MarshalJSON marshals the enum as a quoted json string
func (apinge APINGExceptionCode) MarshalJSON() ([]byte, error) {
	elem, ok := apingExceptionToString[apinge]
	if ok {
		return bytes.NewBufferString(fmt.Sprintf(`"%s"`, elem)).Bytes(), nil
	}

	return bytes.NewBufferString("").Bytes(), errors.New("Yolo error")
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (apinge *APINGExceptionCode) UnmarshalJSON(data []byte) error {
	var j string
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*apinge = apingExceptionToEnum[j]
	return nil
}

// ==================================

// OrderStatus ENUM

type OrderStatus int

const (
	PENDING OrderStatus = iota
	EXECUTION_COMPLETE
	EXECUTABLE
	EXPIRED
)

func (os OrderStatus) String() string {
	return [...]string{"PENDING", "EXECUTION_COMPLETE", "EXECUTABLE", "EXPIRED"}[os]
}

// MarketBettingType ENUM

type MarketBettingType int

const (
	ODDS MarketBettingType = iota
	LINE
	RANGE
	ASIAN_HANDICAP_DOUBLE_LINE
	ASIAN_HANDICAP_SINGLE_LINE
	FIXED_ODDS
)

func (mbt MarketBettingType) String() string {
	return [...]string{"ODDS", "LINE", "RANGE", "ASIAN_HANDICAP_DOUBLE_LINE",
		"ASIAN_HANDICAP_SINGLE_LINE", "FIXED_ODDS"}[mbt]
}
