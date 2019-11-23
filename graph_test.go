package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateGraphCreateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createCreateRequestParameter(
		"name", "times", TypeInt, ColorShibafu, "UTC", SelfSufficientIncrement, true, true)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"graph-id","name":"name","unit":"times","type":"int","color":"shibafu","timezone":"UTC","selfSufficient":"increment","isSecret":true,"publishOptionalData":true}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraphCreate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Create(
		"name", "times", TypeInt, ColorShibafu, "UTC", SelfSufficientIncrement, true, true)

	testSuccess(t, result, err)
}

func TestGraphCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Create(
		"name", "times", TypeInt, ColorShibafu, "UTC", SelfSufficientIncrement, true, true)

	testAPIFailedResult(t, result, err)
}

func TestGraphCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).Create(
		"name", "times", TypeInt, ColorShibafu, "UTC", SelfSufficientIncrement, true, true)

	testPageNotFoundError(t, err)
}

func TestCreateGraphGetAllRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createGetAllRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs", userName)
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

func TestGraphGetAll(t *testing.T) {
	s := `{"graphs":[{"id":"test-graph","name":"graph-name","unit":"commit","type":"int","color":"shibafu","timezone":"Asia/Tokyo","purgeCacheURLs":["https://camo.githubusercontent.com/xxx/xxxx"],"selfSufficient":"increment","isSecret":true,"publishOptionalData":true}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	definitions, err := client.Graph(graphID).GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &GraphDefinitions{
		Graphs: []GraphDefinition{
			{
				ID:                  "test-graph",
				Name:                "graph-name",
				Unit:                "commit",
				Type:                "int",
				Color:               "shibafu",
				TimeZone:            "Asia/Tokyo",
				PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
				SelfSufficient:      "increment",
				IsSecret:            true,
				PublishOptionalData: true,
			},
		},
		Result: Result{IsSuccess: true},
	}
	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestGraphGetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraphGetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).GetAll()

	testPageNotFoundError(t, err)
}

func TestCreateGraphGetSVGRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createGetSVGRequestParameter("20180101", ModeShort)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s?date=20180101&mode=short", userName, graphID)
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

func TestGraphGetSVG(t *testing.T) {
	s := `<svg></svg>`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	svg, err := client.Graph(graphID).GetSVG("20180101", ModeShort)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := s
	if svg != expect {
		t.Errorf("got: %s\nwant: %s", svg, expect)
	}
}

func TestGraphGetSVGFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).GetSVG("20180101", ModeShort)
	expect := "failed to do request: failed to call API: " + string(clientMock.body)
	if err == nil {
		t.Errorf("got: nil\nwant: %s", expect)
	}

	if err != nil && err.Error() != expect {
		t.Errorf("got: %s\nwant: %s", err.Error(), expect)
	}
}

func TestGraphUrl(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	baseURL := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s.html", userName, graphID)
	params := []struct {
		mode   string
		expect string
	}{
		{mode: "", expect: baseURL},
		{mode: "simple", expect: baseURL + "?mode=simple"},
		{mode: "simple-short", expect: baseURL + "?mode=simple-short"},
		{mode: "badge", expect: baseURL + "?mode=badge"},
	}

	for _, p := range params {
		url := client.Graph(graphID).URL(p.mode)
		if url != p.expect {
			t.Errorf("got: %s\nwant: %s", url, p.expect)
		}
	}
}

func TestGraphsUrl(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	url := client.Graph("").GraphsURL()
	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs.html", userName)
	if url != expect {
		t.Errorf("got: %s\nwant: %s", url, expect)
	}
}

func TestCreateStatsRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createStatsRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/stats", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestGraphStats(t *testing.T) {
	s := `{"totalPixelsCount":1,"maxQuantity":2,"minQuantity":3,"totalQuantity":4,"avgQuantity":5.0,"todaysQuantity":6}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	stats, err := client.Graph(graphID).Stats()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Stats{
		TotalPixelsCount: 1,
		MaxQuantity:      2,
		MinQuantity:      3,
		TotalQuantity:    4,
		AvgQuantity:      5.0,
		TodaysQuantity:   6,
		Result:           Result{IsSuccess: true},
	}
	if *stats != *expect {
		t.Errorf("got: %v\nwant: %v", stats, expect)
	}
}

func TestGraphStatsFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Stats()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraphStatsError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).Stats()

	testPageNotFoundError(t, err)
}

func TestCreateGraphUpdateRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createUpdateRequestParameter(
		"name", "times", ColorShibafu, "UTC", []string{"https://camo.githubusercontent.com/xxx/xxxx"}, SelfSufficientIncrement, true, true)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"name","unit":"times","color":"shibafu","timezone":"UTC","purgeCacheURLs":["https://camo.githubusercontent.com/xxx/xxxx"],"selfSufficient":"increment","isSecret":true,"publishOptionalData":true}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraphUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Update(
		"name", "times", ColorShibafu, "UTC", []string{"https://camo.githubusercontent.com/xxx/xxxx"}, SelfSufficientIncrement, true, true)

	testSuccess(t, result, err)
}

func TestGraphUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Update(
		"name", "times", ColorShibafu, "UTC", []string{"https://camo.githubusercontent.com/xxx/xxxx"}, SelfSufficientIncrement, true, true)

	testAPIFailedResult(t, result, err)
}

func TestGraphUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).Update(
		"name", "times", ColorShibafu, "UTC", []string{"https://camo.githubusercontent.com/xxx/xxxx"}, SelfSufficientIncrement, true, true)

	testPageNotFoundError(t, err)
}

func TestCreateGraphDeleteRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createDeleteRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", userName, graphID)
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

func TestGraphDelete(t *testing.T) {
	clientMock = newOKMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Delete()

	testSuccess(t, result, err)
}

func TestGraphDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).Delete()

	testAPIFailedResult(t, result, err)
}

func TestGraphDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).Delete()

	testPageNotFoundError(t, err)
}

func TestCreateGraphGetPixelDatesRequestParameter(t *testing.T) {
	client := Client{UserName: userName, Token: token}
	param, err := client.Graph(graphID).createGetPixelDatesRequestParameter("20180101", "20181231")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/pixels?from=20180101&to=20181231", userName, graphID)
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

func TestGraphGetPixelDates(t *testing.T) {
	s := `{"pixels":["20180101","20180331"]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := Client{UserName: userName, Token: token}
	pixels, err := client.Graph(graphID).GetPixelDates("20180101", "20181231")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Pixels{
		Pixels: []string{"20180101", "20180331"},
		Result: Result{IsSuccess: true},
	}
	if reflect.DeepEqual(pixels, expect) == false {
		t.Errorf("got: %v\nwant: %v", pixels, expect)
	}
}

func TestGraphGetPixelDatesFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := Client{UserName: userName, Token: token}
	result, err := client.Graph(graphID).GetPixelDates("20180101", "20181231")
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraphGetPixelDatesError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := Client{UserName: userName, Token: token}
	_, err := client.Graph(graphID).GetPixelDates("20180101", "20181231")

	testPageNotFoundError(t, err)
}
