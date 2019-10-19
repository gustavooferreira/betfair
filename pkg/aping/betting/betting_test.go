package aping

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {

	bs := BetfairAPI{AppKey: "jI7Ws9goNGfK5gGs", SessionToken: "ocTn3RAhj0PIUpb+LrWTK5dyp9OZhhyzM/s3HG7W0ok="}

	from := time.Date(2019, 8, 21, 0, 0, 0, 0, time.UTC)
	to := time.Date(2019, 8, 21, 23, 55, 0, 0, time.UTC)
	tr := TimeRange{From: &from, To: &to}

	// "sort":"FIRST_TO_START"
	mf := MarketFilter{EventTypeIds: []string{"7"}, MarketCountries: []string{"GB", "IE"},
		MarketTypeCodes: []string{"WIN"}, MarketStartTime: tr}

	bs.ListMarketCatalogue(mf, &[]MarketProjection{MarketProjection_MarketStartTime}, nil, 100)

	if "yolo" != "yolo" {
		t.Errorf("mismatched token: %s", "YOLO")
	}
}

func TestEnumsMarshal(t *testing.T) {
	// "marketProjection":["MARKET_START_TIME"],"sort":"FIRST_TO_START"

	mp := []MarketProjection{MarketProjection_MarketStartTime}
	// mp := MARKET_START_TIME
	t.Logf("Enum: %s", mp)

	bytes, err := json.Marshal(mp)
	if err != nil {
		log.Fatal("error while marshalling")
	}

	if string(bytes) != `["MARKET_START_TIME"]` {
		t.Errorf("mismatched token: %s", string(bytes))
	}
}

func TestEnumsUnmarshal(t *testing.T) {
	message := `["MARKET_START_TIME", "EVENT"]`

	mp := []MarketProjection{}

	err := json.Unmarshal([]byte(message), &mp)
	if err != nil {
		log.Fatal("error while unmarshalling")
	}

	// if string(bytes) != `["MARKET_START_TIME"]` {
	// 	t.Errorf("mismatched token: %s", string(bytes))
	// }
}
