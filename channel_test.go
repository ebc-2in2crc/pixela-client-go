package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateSlackChannelCreateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	param, err := client.Channel().createCreateSlackChannelRequestParameter(
		"channel-id", "channel-name", detail)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"channel-id","name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestSlackChannelCreate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	result, err := client.Channel().CreateSlackChannel("channel-id", "channel-name", detail)

	testSuccess(t, result, err)
}

func TestSlackChannelCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	result, err := client.Channel().CreateSlackChannel("channel-id", "channel-name", detail)

	testAPIFailedResult(t, result, err)
}

func TestSlackChannelCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	_, err := client.Channel().CreateSlackChannel("channel-id", "channel-name", detail)

	testPageNotFoundError(t, err)
}

func TestCreateGetRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Channel().createGetRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels", userName)
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

func TestChannelGetAll(t *testing.T) {
	s := `{"channels":[{"id":"channel-id","name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	definitions, err := client.Channel().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if len(definitions.Channels) != 1 {
		t.Errorf("got: %d definitions \nwant: one", len(definitions.Channels))
	}

	channel := definitions.Channels[0]
	if channel.ID != "channel-id" {
		t.Errorf("got: %v\nwant: %v", channel.ID, "channel-id")
	}
	if channel.Name != "channel-name" {
		t.Errorf("got: %v\nwant: %v", channel.Name, "channel-name")
	}
	if channel.Type != "slack" {
		t.Errorf("got: %v\nwant: %v", channel.Type, "slack")
	}

	d := channel.Detail.(SlackDetail)
	expect := SlackDetail{
		URL:         "slack-url",
		UserName:    "slack-user",
		ChannelName: "slack-channel-name",
	}
	if reflect.DeepEqual(d, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestChannelGetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Channel().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestChannelGetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Channel().GetAll()

	testPageNotFoundError(t, err)
}

func TestCreateSlackChannelUpdateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	param, err := client.Channel().createUpdateSlackChannelRequestParameter(
		"channel-id", "channel-name", detail)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels/channel-id", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestSlackChannelUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	result, err := client.Channel().UpdateSlackChannel("channel-id", "channel-name", detail)

	testSuccess(t, result, err)
}

func TestSlackChannelUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	result, err := client.Channel().UpdateSlackChannel("channel-id", "channel-name", detail)

	testAPIFailedResult(t, result, err)
}

func TestSlackChannelUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	detail := &SlackDetail{URL: "slack-url", UserName: "slack-user", ChannelName: "slack-channel-name"}
	_, err := client.Channel().UpdateSlackChannel("channel-id", "channel-name", detail)

	testPageNotFoundError(t, err)
}

func TestCreateChannelDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Channel().createDeleteSlackChannelRequestParameter("channel-id")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels/channel-id", userName)
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

func TestChannelDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Channel().Delete("channel-id")

	testSuccess(t, result, err)
}

func TestChannelDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Channel().Delete("channel-id")

	testAPIFailedResult(t, result, err)
}

func TestChannelDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Channel().Delete("channel-id")

	testPageNotFoundError(t, err)
}
