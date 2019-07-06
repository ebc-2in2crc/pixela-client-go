package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateUserCreateRequestParameter(t *testing.T) {
	client := Client{UserName: "name", Token: "token"}
	param, err := client.user().createCreateRequestParameter(true, true, "thanks-code")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL + "/users")
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	s := `{"token":"token","username":"name","AgreeTermsOfService":"yes","NotMinor":"yes","thanksCode":"thanks-code"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestUserCreate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.CreateUser(true, true, "thanks-code")

	testSuccess(t, result, err)
}

func TestUserCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.CreateUser(true, true, "thanks-code")

	testAPIFailedResult(t, result, err)
}

func TestUserCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.CreateUser(true, true, "thanks-code")

	testPageNotFoundError(t, err)
}

func TestCreateUserUpdateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.user().createUpdateRequestParameter("newtoken", "thanks-code")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"newToken":"newtoken","thanksCode":"thanks-code"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestUserUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.UpdateUser("newToken", "thanks-code")

	testSuccess(t, result, err)

	if client.Token != "newToken" {
		t.Errorf("got: %s\nwant: newToken", client.Token)
	}
}

func TestUserUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.UpdateUser("newToken", "thanks-code")

	testAPIFailedResult(t, result, err)

	if client.Token != token {
		t.Errorf("got: %s\nwant: %s", client.Token, token)
	}
}

func TestUserUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.UpdateUser("newToken", "thanks-code")

	testPageNotFoundError(t, err)

	if client.Token != token {
		t.Errorf("got: %s\nwant: %s", client.Token, token)
	}
}

func TestCreateUserDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.user().createDeleteRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s", userName)
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

func TestUserDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.DeleteUser()

	testSuccess(t, result, err)
}

func TestUserDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.DeleteUser()

	testAPIFailedResult(t, result, err)
}

func TestUserDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.DeleteUser()

	testPageNotFoundError(t, err)
}
