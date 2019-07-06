package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateWebhookCreateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Webhook().createCreateRequestParameter(graphID, SelfSufficientIncrement)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"graphID":"graph-id","type":"increment"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestWebhookCreate(t *testing.T) {
	s := `{"webhookHash":"webhook-hash","message":"Success.","isSuccess":true}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Create(graphID, SelfSufficientIncrement)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &WebhookCreateResult{
		WebhookHash: "webhook-hash",
		Result:      Result{Message: "Success.", IsSuccess: true},
	}
	if *result != *expect {
		t.Errorf("got: %v\nwant: %v", result, expect)
	}
}

func TestWebhookCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Create(graphID, SelfSufficientIncrement)

	testAPIFailedResult(t, &result.Result, err)
}

func TestWebhookCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Webhook().Create(graphID, SelfSufficientIncrement)

	testPageNotFoundError(t, err)
}

func TestCreateWebhookGetAllRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Webhook().createGetAllRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestWebhookGetAll(t *testing.T) {
	s := `{"webhooks":[{"webhookHash":"webhook-hash","graphID":"test-graph","type":"increment"}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	definitions, err := client.Webhook().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &WebhookDefinitions{
		Webhooks: []WebhookDefinition{
			{
				WebhookHash: "webhook-hash",
				GraphID:     "test-graph",
				Type:        "increment",
			},
		},
		Result: Result{IsSuccess: true},
	}
	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestWebhookGetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().GetAll()

	testAPIFailedResult(t, &result.Result, err)
}

func TestWebhookGetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Webhook().GetAll()

	testPageNotFoundError(t, err)
}

func TestCreateWebhookDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Webhook().createDeleteRequestParameter("webhook-hash")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/webhook-hash", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestWebhookDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Delete("webhook-hash")

	testSuccess(t, result, err)
}

func TestWebhookDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Delete("webhook-hash")

	testAPIFailedResult(t, result, err)
}

func TestCreateWebhookInvokeRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Webhook().createInvokeRequestParameter("webhook-hash")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/webhook-hash", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestWebhookInvoke(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Invoke("webhook-hash")

	testSuccess(t, result, err)
}

func TestWebhookInvokeFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Webhook().Invoke("webhook-hash")

	testAPIFailedResult(t, result, err)
}

func TestWebhookInvokeError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Webhook().Invoke("webhook-hash")

	testPageNotFoundError(t, err)
}
