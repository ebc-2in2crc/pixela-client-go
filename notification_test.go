package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateNotificationCreateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Notification(graphID).createCreateNotificationRequestParameter(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"notification-id","name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","channelID":"channel-id"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestNotificationCreate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Create(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testSuccess(t, result, err)
}

func TestNotificationCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Create(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testAPIFailedResult(t, result, err)
}

func TestNotificationCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Notification(graphID).Create(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testPageNotFoundError(t, err)
}

func TestCreateNotificationGetRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Notification(graphID).createGetNotificationRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", userName, graphID)
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

func TestNotificationGetAll(t *testing.T) {
	s := `{"notifications":[{"id":"notification-id","name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","channelID":"channel-id"}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	definitions, err := client.Notification(graphID).GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &NotificationDefinitions{
		Notifications: []NotificationDefinition{
			{
				ID:        "notification-id",
				Name:      "notification-name",
				Target:    TargetQuantity,
				Condition: ConditionGreaterThan,
				Threshold: "3",
				ChannelID: "channel-id",
			},
		},
		Result: Result{IsSuccess: true},
	}

	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestCreateNotificationUpdateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Notification(graphID).createUpdateNotificationRequestParameter(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", userName, graphID, "notification-id")
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","channelID":"channel-id"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestNotificationUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Update(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testSuccess(t, result, err)
}

func TestNotificationUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Update(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testAPIFailedResult(t, result, err)
}

func TestNotificationUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Notification(graphID).Update(
		"notification-id", "notification-name", TargetQuantity, ConditionGreaterThan, "3", "channel-id")

	testPageNotFoundError(t, err)
}

func TestCreateNotificationDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Notification(graphID).createDeleteNotificationRequestParameter("notification-id")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", userName, graphID, "notification-id")
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

func TestNotificationDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Delete("notification-id")

	testSuccess(t, result, err)
}

func TestNotificationDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Notification(graphID).Delete("notification-id")

	testAPIFailedResult(t, result, err)
}

func TestNotificationDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Notification(graphID).Delete("notification-id")

	testPageNotFoundError(t, err)
}
