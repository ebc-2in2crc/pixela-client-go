package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Webhook manages communication with the Pixela webhook API.
type Webhook struct {
	UserName string
	Token    string
}

// Create create a new Webhook.
func (w *Webhook) Create(graphID, selfSufficient string) (*WebhookCreateResult, error) {
	param, err := w.createCreateRequestParameter(graphID, selfSufficient)
	if err != nil {
		return &WebhookCreateResult{}, errors.Wrapf(err, "failed to create webhook create parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &WebhookCreateResult{}, errors.Wrapf(err, "failed to do request")
	}

	var createResult WebhookCreateResult
	if err := json.Unmarshal(b, &createResult); err != nil {
		return &WebhookCreateResult{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	return &createResult, nil
}

// WebhookCreateResult is Create() Result struct.
type WebhookCreateResult struct {
	WebhookHash string `json:"webhookHash"`
	Result
}

func (w *Webhook) createCreateRequestParameter(graphID, selfSufficient string) (*requestParameter, error) {
	create := webhookCreate{GraphID: graphID, Type: selfSufficient}
	b, err := json.Marshal(&create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", w.UserName),
		Header: map[string]string{userToken: w.Token},
		Body:   b,
	}, nil
}

type webhookCreate struct {
	GraphID string `json:"graphID"`
	Type    string `json:"type"`
}

// GetAll get all predefined webhooks definitions.
func (w *Webhook) GetAll() (*WebhookDefinitions, error) {
	param, err := w.createGetAllRequestParameter()
	if err != nil {
		return &WebhookDefinitions{}, errors.Wrapf(err, "failed to create get all webhooks parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &WebhookDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions WebhookDefinitions
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &WebhookDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

// WebhookDefinitions is webhook definition list.
type WebhookDefinitions struct {
	Webhooks []WebhookDefinition `json:"webhooks"`
	Result
}

// WebhookDefinition is webhook definition.
type WebhookDefinition struct {
	WebhookHash string `json:"webhookHash"`
	GraphID     string `json:"graphId"`
	Type        string `json:"type"`
}

func (w *Webhook) createGetAllRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", w.UserName),
		Header: map[string]string{userToken: w.Token},
		Body:   []byte{},
	}, nil
}

// Delete delete the registered Webhook.
func (w *Webhook) Delete(webhookHash string) (*Result, error) {
	param, err := w.createDeleteRequestParameter(webhookHash)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create webhook delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (w *Webhook) createDeleteRequestParameter(webhookHash string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/%s", w.UserName, webhookHash),
		Header: map[string]string{userToken: w.Token},
		Body:   []byte{},
	}, nil
}

// Invoke invoke the webhook registered in advance.
// It is used "timezone" setting as post date if Graph's "timezone" is specified, if not specified, calculates it in "UTC".
func (w *Webhook) Invoke(webhookHash string) (*Result, error) {
	param, err := w.createInvokeRequestParameter(webhookHash)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create webhook invoke parameter")
	}

	return doRequestAndParseResponse(param)
}

func (w *Webhook) createInvokeRequestParameter(webhookHash string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/%s", w.UserName, webhookHash),
		Header: map[string]string{contentLength: "0"},
		Body:   []byte{},
	}, nil
}
