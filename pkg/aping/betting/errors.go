package betting

import "fmt"

type BettingAPIError struct {
	ErrorCode    APINGExceptionCode
	ErrorDetails string
	RequestUUID  string
}

func (e *BettingAPIError) Error() string {
	return fmt.Sprintf("Betfair APING error: %s - Details: %s - RequestUUID: %s", e.ErrorCode, e.ErrorDetails, e.RequestUUID)
}
