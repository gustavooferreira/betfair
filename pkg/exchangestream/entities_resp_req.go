package exchangestream

import (
	"encoding/json"
	"errors"
)

type RequestMessage struct {
	Op                        string  `json:"op"`
	ID                        *uint32 `json:"id,omitempty"`
	AuthenticationMessage     *AuthenticationMessage
	MarketSubscriptionMessage *MarketSubscriptionMessage
	OrderSubscriptionMessage  *OrderSubscriptionMessage
}

// MarshalJSON marshals RequestMessage struct
func (rm RequestMessage) MarshalJSON() ([]byte, error) {
	if rm.Op == "heartbeat" {
		return json.Marshal(&struct {
			Op string  `json:"op"`
			ID *uint32 `json:"id,omitempty"`
		}{
			Op: rm.Op,
			ID: rm.ID,
		})
	} else if rm.Op == "authentication" && rm.AuthenticationMessage != nil {
		return json.Marshal(&struct {
			Op           string  `json:"op"`
			ID           *uint32 `json:"id,omitempty"`
			AppKey       string  `json:"appKey"`
			SessionToken string  `json:"session"`
		}{
			Op:           rm.Op,
			ID:           rm.ID,
			AppKey:       rm.AuthenticationMessage.AppKey,
			SessionToken: rm.AuthenticationMessage.SessionToken,
		})
	} else if rm.Op == "marketSubscription" && rm.MarketSubscriptionMessage != nil {
		return json.Marshal(&struct {
			Op                  string           `json:"op"`
			ID                  *uint32          `json:"id,omitempty"`
			SegmentationEnabled bool             `json:"segmentationEnabled"`
			Clk                 string           `json:"clk,omitempty"`
			HeartbeatMs         uint             `json:"heartbeatMs,omitempty"`
			InitialClk          string           `json:"initialClk,omitempty"`
			MarketFilter        MarketFilter     `json:"marketFilter"`
			ConflateMs          uint             `json:"conflateMs,omitempty"`
			MarketDataFilter    MarketDataFilter `json:"marketDataFilter"`
		}{
			Op:                  rm.Op,
			ID:                  rm.ID,
			SegmentationEnabled: rm.MarketSubscriptionMessage.SegmentationEnabled,
			Clk:                 rm.MarketSubscriptionMessage.Clk,
			HeartbeatMs:         rm.MarketSubscriptionMessage.HeartbeatMs,
			InitialClk:          rm.MarketSubscriptionMessage.InitialClk,
			MarketFilter:        rm.MarketSubscriptionMessage.MarketFilter,
			ConflateMs:          rm.MarketSubscriptionMessage.ConflateMs,
			MarketDataFilter:    rm.MarketSubscriptionMessage.MarketDataFilter,
		})
	} else if rm.Op == "orderSubscription" && rm.OrderSubscriptionMessage != nil {
		return json.Marshal(&struct {
			Op                  string      `json:"op"`
			ID                  *uint32     `json:"id,omitempty"`
			SegmentationEnabled bool        `json:"segmentationEnabled"`
			OrderFilter         OrderFilter `json:"orderFilter"`
			Clk                 string      `json:"clk,omitempty"`
			HeartbeatMs         uint        `json:"heartbeatMs,omitempty"`
			InitialClk          string      `json:"initialClk,omitempty"`
			ConflateMs          uint        `json:"conflateMs,omitempty"`
		}{
			Op:                  rm.Op,
			ID:                  rm.ID,
			SegmentationEnabled: rm.OrderSubscriptionMessage.SegmentationEnabled,
			OrderFilter:         rm.OrderSubscriptionMessage.OrderFilter,
			Clk:                 rm.OrderSubscriptionMessage.Clk,
			HeartbeatMs:         rm.OrderSubscriptionMessage.HeartbeatMs,
			InitialClk:          rm.OrderSubscriptionMessage.InitialClk,
			ConflateMs:          rm.OrderSubscriptionMessage.ConflateMs,
		})
	}

	return json.Marshal(nil)
}

type ResponseMessage struct {
	Op                  string  `json:"op"`
	ID                  *uint32 `json:"id,omitempty"`
	ConnectionMessage   *ConnectionMessage
	StatusMessage       *StatusMessage
	MarketChangeMessage *MarketChangeMessage
	OrderChangeMessage  *OrderChangeMessage
}

// UnmarshalJSON unmarshals ResponseMessage struct
func (rm *ResponseMessage) UnmarshalJSON(data []byte) error {
	temp := struct {
		Op string  `json:"op"`
		ID *uint32 `json:"id,omitempty"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	rm.Op = temp.Op
	rm.ID = temp.ID

	if rm.Op == "connection" {
		var connectionMessage ConnectionMessage
		if err := json.Unmarshal(data, &connectionMessage); err != nil {
			return err
		}
		rm.ConnectionMessage = &connectionMessage
	} else if rm.Op == "status" {
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
	} else if rm.Op == "ocm" {
		var orderChangeMessage OrderChangeMessage
		if err := json.Unmarshal(data, &orderChangeMessage); err != nil {
			return err
		}
		rm.OrderChangeMessage = &orderChangeMessage
	} else {
		return errors.New("Invalid object value")
	}

	return nil
}
