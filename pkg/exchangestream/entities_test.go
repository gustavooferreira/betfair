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

// // Construct Request
// authMsg := exchangestream.AuthenticationMessage{AppKey: "app_key", SessionToken: "session_token"}
// reqMsg := exchangestream.RequestMessage{Op: "authentication", ID: 100, AuthenticationMessage: &authMsg}
// bytes, err := json.Marshal(reqMsg)
// if err != nil {
// 	fmt.Printf("ERROR: %T - %+[1]v\n", err)
// 	return
// }
// fmt.Printf("%s\n", bytes)

// // Deserialize Response
// respMsgString := `{"op":"status", "id":10, "errorMessage":"yolo error message", "errorCode": "UNEXPECTED_ERROR", "connectionId": "123-321", "connectionClosed": true, "statusCode": "SUCCESS"}`

// respMsg := exchangestream.ResponseMessage{}
// err = json.Unmarshal([]byte(respMsgString), &respMsg)
// if err != nil {
// 	fmt.Printf("ERROR: %T - %+[1]v\n", err)
// 	return
// }
// fmt.Printf("%+v\n", respMsg)
// fmt.Printf("%+v\n", respMsg.StatusMessage)
// fmt.Printf("Connection closed: %+v\n", *respMsg.StatusMessage.ConnectionClosed)

// if respMsg.StatusMessage.ErrorCode == exchangestream.ErrorCode_UnexpectedError {
// 	fmt.Println("YOLOOOOO")
// }
