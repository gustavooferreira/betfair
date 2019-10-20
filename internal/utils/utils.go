package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type BetfairAPIError struct {
	StatusCode int
	Body       string
}

func (e *BetfairAPIError) Error() string {
	return fmt.Sprintf("unexpected status code: %d - body: %s", e.StatusCode, e.Body)
}

// SendRequest sends request to the betfair servers
// TODO: log requests and responses (headers, body, etc)
func SendRequest(httpClient *http.Client, method string, appKey string, sessionToken string, url string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if appKey != "" {
		req.Header.Set("X-Application", appKey)
	}

	if sessionToken != "" {
		req.Header.Set("X-Authentication", sessionToken)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &BetfairAPIError{StatusCode: resp.StatusCode, Body: string(buf)}
	}

	return buf, nil
}
