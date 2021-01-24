package auth_test

import (
	"testing"

	"github.com/gustavooferreira/betfair/auth"
)

func TestAuth(t *testing.T) {

	lr := auth.LoginResponse{LoginStatus: "SUCCESS", SessionToken: "myToken"}

	if lr.SessionToken != "myToken" {
		t.Errorf("mismatched token: %s", lr.SessionToken)
	}
}
