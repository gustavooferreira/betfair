package exchangestream

import "fmt"

// ConnectionError is the error returned by Connect() when an error occurs
type ConnectionError struct {
	Msg string
	Err error
}

func (e ConnectionError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s", e.Msg, e.Err)
	}
	return fmt.Sprintf("%s", e.Msg)
}

func (e ConnectionError) Unwrap() error {
	return e.Err
}
