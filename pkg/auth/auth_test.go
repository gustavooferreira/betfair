package auth

import "testing"

func TestAbs(t *testing.T) {

	lr := LoginResponse{LoginStatus: "SUCCESS", SessionToken: "myToken"}

	if lr.SessionToken != "myToken" {
		t.Errorf("mismatched token: %s", lr.SessionToken)
	}
}
