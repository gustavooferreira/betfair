package exchangestream

import (
	"encoding/json"
	"testing"
)

func TestConnectionMessage(t *testing.T) {
	msg := `{"op":"connection", "connectionId":"123-321"}`

	msgStruct := ResponseMessage{}
	err := json.Unmarshal([]byte(msg), &msgStruct)
	if err != nil {
		t.Errorf("error while decoding message - error: %s", err)
	}

	if msgStruct.Op != "connection" || msgStruct.ConnectionMessage.ConnectionID != "123-321" {
		t.Errorf("Message not properly decoded, got: %+v, want: %s.", msgStruct, msg)
	}
}
