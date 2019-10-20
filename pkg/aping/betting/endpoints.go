package betting

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/gustavooferreira/betfair/internal/utils"
	"github.com/gustavooferreira/betfair/pkg/aping"
)

const ukBettingEndpoint = "https://api.betfair.com/exchange/betting/rest/v1.0/"

const (
	listMarketCatalogueEndpoint = ukBettingEndpoint + "listMarketCatalogue/"
	listMarketBookEndpoint      = ukBettingEndpoint + "listMarketBook/"
	placeOrdersEndpoint         = ukBettingEndpoint + "placeOrders/"
)

type BettingAPIError struct {
	ErrorCode    APINGExceptionCode
	ErrorDetails string
	RequestUUID  string
}

func (e *BettingAPIError) Error() string {
	return fmt.Sprintf("Betfair APING error: %s - Details: %s - RequestUUID: %s", e.ErrorCode, e.ErrorDetails, e.RequestUUID)
}

type BettingAPI struct {
	aping.BetfairAPI
}

func NewBettingAPI(bapi aping.BetfairAPI) BettingAPI {
	bettingAPI := BettingAPI{bapi}
	return bettingAPI
}

// ListMarketCatalogue lists the market catalogue.
// Note: listMarketCatalogue does not return markets that are CLOSED.
// Only filter and maxResults are mandatory.
// All the other arguments are optional and therefore they are pointers so the user can pass nil in case they don't want to set them
func (b BettingAPI) ListMarketCatalogue(filter MarketFilter, mp *[]MarketProjection, marketSort *MarketSort, maxResults uint, locale *string) ([]MarketCatalogue, error) {
	lrc := listMarketCatalogueReqContainer{Filter: filter, MarketProjection: mp, Sort: marketSort, MaxResults: maxResults, Locale: locale}

	lrcBytes, err := json.Marshal(lrc)
	if err != nil {
		log.Fatal("error while marshalling request")
	}

	payload := bytes.NewBuffer(lrcBytes)
	response, err := b.sendRequest(listMarketCatalogueEndpoint, payload)
	if err != nil {
		return nil, err
	}

	mcs := []MarketCatalogue{}

	err = json.Unmarshal(response, &mcs)
	if err != nil {
		return []MarketCatalogue{}, errors.New("error while unmarshalling response")
	}

	return mcs, nil
}

// ListMarketBook lists dynamic data about markets.
func (b BettingAPI) ListMarketBook(marketIDs []string) ([]MarketBook, error) {
	lrc := listMarketBookReqContainer{MarketIDs: marketIDs}

	lrcBytes, err := json.Marshal(lrc)
	if err != nil {
		log.Fatal("error while marshalling request")
	}

	payload := bytes.NewBuffer(lrcBytes)
	response, err := b.sendRequest(listMarketBookEndpoint, payload)
	if err != nil {
		return nil, err
	}

	mbs := []MarketBook{}

	err = json.Unmarshal(response, &mbs)
	if err != nil {
		return []MarketBook{}, errors.New("error while unmarshalling response")
	}

	return mbs, nil
}

// PlaceOrders puts back/lay bets on the market
func (b BettingAPI) PlaceOrders(marketID string, instructions []PlaceInstruction) (PlaceExecutionReport, error) {
	prc := placeOrderReqContainer{MarketID: marketID, Instructions: instructions}

	prcBytes, err := json.Marshal(prc)
	if err != nil {
		log.Fatal("error while marshalling request")
	}

	payload := bytes.NewBuffer(prcBytes)
	response, err := b.sendRequest(placeOrdersEndpoint, payload)
	if err != nil {
		return PlaceExecutionReport{}, err
	}

	per := PlaceExecutionReport{}

	err = json.Unmarshal(response, &per)
	if err != nil {
		return PlaceExecutionReport{}, errors.New("error while unmarshalling response")
	}

	return per, nil
}

func (b BettingAPI) sendRequest(url string, body io.Reader) ([]byte, error) {

	respBody, err := utils.SendRequest(b.HttpClient, "POST", b.AppKey, b.SessionToken, url, body)

	// Encapsulate error here!
	if errB, ok := err.(*utils.BetfairAPIError); ok {
		bapie := BetfairAPIError{}
		err = json.Unmarshal([]byte(errB.Body), &bapie)
		if err != nil {
			return nil, errors.New("error while unmarshalling APINGException response")
		}

		return nil, &BettingAPIError{ErrorCode: bapie.Detail.APINGException.ErrorCode,
			ErrorDetails: bapie.Detail.APINGException.ErrorDetails, RequestUUID: bapie.Detail.APINGException.RequestUUID}
	} else if err != nil {
		return nil, err
	}

	return respBody, nil
}
