package aping

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const ukBettingEndpoint = "https://api.betfair.com/rest/v1.0/"

const (
	listEventTypesEndpoint      = ukBettingEndpoint + "listEventTypes/"
	listMarketCatalogueEndpoint = ukBettingEndpoint + "listMarketCatalogue/"
)

var ukEndpoints = map[string]string{
	"account": "https://api.betfair.com/exchange/account/rest/v1.0/",
}

type BetfairAPINGError struct {
	ErrorCode    APINGExceptionCode
	ErrorDetails string
}

func (e *BetfairAPINGError) Error() string {
	return fmt.Sprintf("Betfair APING error: %s - Details: %s", e.ErrorCode, e.ErrorDetails)
}

type BetfairAPI struct {
	httpClient   *http.Client
	AppKey       string
	SessionToken string
}

func NewBetfairAPI(httpClient *http.Client, appKey string, sessionToken string) BetfairAPI {
	bapi := BetfairAPI{httpClient: httpClient, AppKey: appKey, SessionToken: sessionToken}
	return bapi
}

// ListMarketCatalogue lists the market catalogue
// Only marketFilter and maxResults is mandatory
// All the other arguments are optional and therefore they are pointers so the user can pass nil in case they don't want to set them
func (b BetfairAPI) ListMarketCatalogue(marketFilter MarketFilter, mp *[]MarketProjection, marketSort *MarketSort, maxResults uint) ([]MarketCatalogue, error) {
	lrc := listMarketCatalogueReqContainer{Filter: marketFilter, MarketProjection: mp,
		Sort: marketSort, MaxResults: maxResults}

	lrcBytes, err := json.Marshal(lrc)
	if err != nil {
		log.Fatal("error while marshalling")
	}

	payload := bytes.NewBuffer(lrcBytes)

	log.Printf("Request body: %s", payload)

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

func (b BetfairAPI) sendRequest(url string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest("POST", listMarketCatalogueEndpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Application", b.AppKey)
	req.Header.Set("X-Authentication", b.SessionToken)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := b.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Encapsulate error here!
	if resp.StatusCode != 200 {

		bapie := BetfairAPIError{}
		err = json.Unmarshal(buf, &bapie)
		if err != nil {
			return nil, errors.New("error while unmarshalling response error response - make this better")
		}

		log.Printf("%+v\n", bapie)

		if bapie.Detail.APINGException.ErrorCode == APINGExceptionCode_InvalidAppKey {
			log.Println("YOOOOOOLOOOOOO")
		}

		// return nil, errors.New(resp.Status)
		return nil, &BetfairAPINGError{ErrorCode: bapie.Detail.APINGException.ErrorCode,
			ErrorDetails: bapie.Detail.APINGException.ErrorDetails}
	}

	return buf, nil
}
