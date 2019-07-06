package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestCreatePixelCreateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createCreateRequestParameter("20180915", "5")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"date":"20180915","quantity":"5"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixelCreate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Create("20180915", "5")

	testSuccess(t, result, err)
}

func TestPixelCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Create("20180915", "5")

	testAPIFailedResult(t, result, err)
}

func TestPixelCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Create("20180915", "5")

	testPageNotFoundError(t, err)
}

func TestCreatePixelIncrementRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createIncrementRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/increment", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}
	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestPixelIncrement(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Increment()

	testSuccess(t, result, err)
}

func TestPixelIncrementFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Increment()

	testAPIFailedResult(t, result, err)
}

func TestPixelIncrementError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Increment()

	testPageNotFoundError(t, err)
}

func TestCreatePixelDecrementRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createDecrementRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/decrement", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}
	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}
}

func TestPixelDecrement(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Decrement()

	testSuccess(t, result, err)
}

func TestPixelDecrementFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Decrement()

	testAPIFailedResult(t, result, err)
}

func TestPixelDecrementError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Decrement()

	testPageNotFoundError(t, err)
}

func TestCreatePixelGetRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createGetRequestParameter("20180915")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/%s", userName, graphID, "20180915")
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

func TestPixelGet(t *testing.T) {
	s := `{"quantity": "5"}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	quantity, err := client.Pixel(graphID).Get("20180915")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Quantity{
		Quantity: "5",
		Result:   Result{IsSuccess: true},
	}
	if *quantity != *expect {
		t.Errorf("got: %v\nwant: %v", quantity, expect)
	}
}

func TestPixelGetFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Get("20180915")

	testAPIFailedResult(t, &result.Result, err)
}

func TestPixelGetError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Get("20180915")

	testPageNotFoundError(t, err)
}

func TestCreatePixelUpdateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createUpdateRequestParameter("20180915", "5")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/20180915", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"5"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixelUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Update("20180915", "5")

	testSuccess(t, result, err)
}

func TestPixelUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Update("20180915", "5")

	testAPIFailedResult(t, result, err)
}

func TestPixelUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Update("20180915", "5")

	testPageNotFoundError(t, err)
}

func TestCreatePixelDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Pixel(graphID).createDeleteRequestParameter("20180915")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/20180915", userName, graphID)
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

func TestPixelDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Delete("20180915")

	testSuccess(t, result, err)
}

func TestPixelDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Pixel(graphID).Delete("20180915")

	testAPIFailedResult(t, result, err)
}

func TestPixelDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Pixel(graphID).Delete("20180915")

	testPageNotFoundError(t, err)
}
