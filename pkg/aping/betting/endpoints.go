package betting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gustavooferreira/betfair/internal/utils"
	"github.com/gustavooferreira/betfair/pkg/aping"
)

const ukBettingEndpoint = "https://api.betfair.com/exchange/betting/rest/v1.0/"

const (
	listMarketCatalogueEndpoint = ukBettingEndpoint + "listMarketCatalogue/"
	listMarketBookEndpoint      = ukBettingEndpoint + "listMarketBook/"
	placeOrdersEndpoint         = ukBettingEndpoint + "placeOrders/"
	replaceOrdersEndpoint       = ukBettingEndpoint + "replaceOrders/"
	cancelOrdersEndpoint        = ukBettingEndpoint + "cancelOrders/"
)

type BettingAPI struct {
	aping.BetfairAPI
}

func NewBettingAPI(bapi aping.BetfairAPI) BettingAPI {
	bettingAPI := BettingAPI{bapi}
	return bettingAPI
}

// ListMarketCatalogue lists the market catalogue.
// Note: listMarketCatalogue does not return markets that are CLOSED.
func (b BettingAPI) ListMarketCatalogue(lrc ContainerListMarketCatalogue) ([]MarketCatalogue, error) {
	lrcBytes, err := json.Marshal(lrc)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling request %w", err)
	}

	payload := bytes.NewBuffer(lrcBytes)
	response, err := b.sendRequest(listMarketCatalogueEndpoint, payload)
	if err != nil {
		return nil, err
	}

	mcs := []MarketCatalogue{}
	err = json.Unmarshal(response, &mcs)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response %w", err)
	}

	return mcs, nil
}

// ListMarketBook lists dynamic data about markets.
// Calls to listMarketBook should be made up to a maximum of 5 times per second to a single marketId.
func (b BettingAPI) ListMarketBook(clmb ContainerListMarketBook) ([]MarketBook, error) {
	clmbBytes, err := json.Marshal(clmb)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling request %w", err)
	}

	payload := bytes.NewBuffer(clmbBytes)
	response, err := b.sendRequest(listMarketBookEndpoint, payload)
	if err != nil {
		return nil, err
	}

	mbs := []MarketBook{}
	err = json.Unmarshal(response, &mbs)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response %w", err)
	}

	return mbs, nil
}

// PlaceOrders puts back/lay bets on the market.
func (b BettingAPI) PlaceOrders(cpo ContainerPlaceOrders) (PlaceExecutionReport, error) {
	per := PlaceExecutionReport{}

	cpoBytes, err := json.Marshal(cpo)
	if err != nil {
		return PlaceExecutionReport{}, fmt.Errorf("error while marshalling request %w", err)
	}

	payload := bytes.NewBuffer(cpoBytes)
	response, err := b.sendRequest(placeOrdersEndpoint, payload)
	if err != nil {
		return per, err
	}

	err = json.Unmarshal(response, &per)
	if err != nil {
		return per, fmt.Errorf("error while unmarshalling response %w", err)
	}

	return per, nil
}

// ReplaceOrders cancels bets followed by putting new bets on the market.
func (b BettingAPI) ReplaceOrders(cro ContainerReplaceOrders) (ReplaceExecutionReport, error) {
	rer := ReplaceExecutionReport{}

	croBytes, err := json.Marshal(cro)
	if err != nil {
		return rer, fmt.Errorf("error while marshalling request %w", err)
	}

	payload := bytes.NewBuffer(croBytes)
	response, err := b.sendRequest(replaceOrdersEndpoint, payload)
	if err != nil {
		return rer, err
	}

	err = json.Unmarshal(response, &rer)
	if err != nil {
		return ReplaceExecutionReport{}, fmt.Errorf("error while unmarshalling response %w", err)
	}

	return rer, nil
}

// CancelOrders cancels bets on the market.
func (b BettingAPI) CancelOrders(cco ContainerCancelOrders) (CancelExecutionReport, error) {
	cer := CancelExecutionReport{}

	ccoBytes, err := json.Marshal(cco)
	if err != nil {
		return cer, fmt.Errorf("error while marshalling request %w", err)
	}

	payload := bytes.NewBuffer(ccoBytes)
	response, err := b.sendRequest(cancelOrdersEndpoint, payload)
	if err != nil {
		return cer, err
	}

	err = json.Unmarshal(response, &cer)
	if err != nil {
		return cer, fmt.Errorf("error while unmarshalling response %w", err)
	}

	return cer, nil
}

func (b BettingAPI) sendRequest(url string, body io.Reader) ([]byte, error) {

	respBody, err := utils.SendRequest(b.HttpClient, "POST", b.AppKey, b.SessionToken, url, body)

	// Encapsulate error here!
	if errB, ok := err.(*utils.BetfairAPIError); ok {
		bapie := BetfairAPIError{}
		err = json.Unmarshal([]byte(errB.Body), &bapie)
		if err != nil {
			return nil, fmt.Errorf("error while unmarshalling APINGException response %w", err)
		}

		return nil, &BettingAPIError{
			ErrorCode:    bapie.Detail.APINGException.ErrorCode,
			ErrorDetails: bapie.Detail.APINGException.ErrorDetails,
			RequestUUID:  bapie.Detail.APINGException.RequestUUID,
		}
	} else if err != nil {
		return nil, err
	}

	return respBody, nil
}
