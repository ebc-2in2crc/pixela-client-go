package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Pixel manages communication with the Pixela pixel API.
type Pixel struct {
	UserName string
	Token    string
	GraphID  string
}

// Create records the quantity of the specified date as a "Pixel".
func (p *Pixel) Create(date string, quantity, optionalData string) (*Result, error) {
	param, err := p.createCreateRequestParameter(date, quantity, optionalData)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel create parameter")
	}

	return doRequestAndParseResponse(param)
}

func (p *Pixel) createCreateRequestParameter(date, quantity, optionalData string) (*requestParameter, error) {
	create := pixelCreate{Date: date, Quantity: quantity, OptionalData: optionalData}
	b, err := json.Marshal(&create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", p.UserName, p.GraphID),
		Header: map[string]string{userToken: p.Token},
		Body:   b,
	}, nil
}

type pixelCreate struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

// Increment increments quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then 1 added, and for float then 0.01 added.
func (p *Pixel) Increment() (*Result, error) {
	param, err := p.createIncrementRequestParameter()
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel increment parameter")
	}

	return doRequestAndParseResponse(param)
}

func (p *Pixel) createIncrementRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/increment", p.UserName, p.GraphID),
		Header: map[string]string{contentLength: "0", userToken: p.Token},
		Body:   []byte{},
	}, nil
}

// Decrement decrements quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then -1 added, and for float then -0.01 added.
func (p *Pixel) Decrement() (*Result, error) {
	param, err := p.createDecrementRequestParameter()
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel decrement parameter")
	}

	return doRequestAndParseResponse(param)
}

func (p *Pixel) createDecrementRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/decrement", p.UserName, p.GraphID),
		Header: map[string]string{contentLength: "0", userToken: token},
		Body:   []byte{},
	}, nil
}

// Get gets registered quantity as "Pixel".
func (p *Pixel) Get(date string) (*Quantity, error) {
	param, err := p.createGetRequestParameter(date)
	if err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to create pixel get parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to do request")
	}

	var quantity Quantity
	if err := json.Unmarshal(b, &quantity); err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	quantity.IsSuccess = quantity.Message == ""
	return &quantity, nil
}

func (p *Pixel) createGetRequestParameter(date string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/%s", p.UserName, p.GraphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   []byte{},
	}, nil
}

// Quantity ... registered quantity.
type Quantity struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
	Result
}

// Update updates the quantity already registered as a "Pixel".
func (p *Pixel) Update(date, quantity, optionalData string) (*Result, error) {
	param, err := p.createUpdateRequestParameter(date, quantity, optionalData)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel update parameter")
	}

	return doRequestAndParseResponse(param)
}

func (p *Pixel) createUpdateRequestParameter(date, quantity, optionalData string) (*requestParameter, error) {
	update := pixelUpdate{Quantity: quantity, OptionalData: optionalData}
	b, err := json.Marshal(update)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/%s", p.UserName, p.GraphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   b,
	}, nil
}

type pixelUpdate struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

// Delete deletes the registered "Pixel".
func (p *Pixel) Delete(date string) (*Result, error) {
	param, err := p.createDeleteRequestParameter(date)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (p *Pixel) createDeleteRequestParameter(date string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/%s", p.UserName, p.GraphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   []byte{},
	}, nil
}
