package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Channel manages communication with the Pixela graph API.
type Channel struct {
	UserName string
	Token    string
}

const (
	slackChannel = "slack"
)

// CreateSlackChannel creates a new slack channel.
func (c *Channel) CreateSlackChannel(channelID, channelName string, detail *SlackDetail) (*Result, error) {
	param, err := c.createCreateSlackChannelRequestParameter(channelID, channelName, detail)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel create parameter")
	}

	return doRequestAndParseResponse(param)
}

func (c *Channel) createCreateSlackChannelRequestParameter(id, name string, detail *SlackDetail) (*requestParameter, error) {
	create := &slackChannelCreate{
		ID:     id,
		Name:   name,
		Type:   slackChannel,
		Detail: detail,
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels", c.UserName),
		Header: map[string]string{userToken: c.Token},
		Body:   b,
	}, nil
}

// SlackDetail is channel detail settings when type is slack.
type SlackDetail struct {
	URL         string `json:"url"`
	UserName    string `json:"userName"`
	ChannelName string `json:"channelName"`
}

type slackChannelCreate struct {
	ID     string       `json:"id,omitempty"`
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Detail *SlackDetail `json:"detail"`
}

// GetAll gets all predefined channels.
func (c *Channel) GetAll() (*ChannelDefinitions, error) {
	param, err := c.createGetRequestParameter()
	if err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to create get all channels parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var raw rawChannelDefinitions
	if err := json.Unmarshal(b, &raw); err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	var definitions ChannelDefinitions
	definitions.Channels = make([]ChannelDefinition, len(raw.Channels))
	for i, v := range raw.Channels {
		definitions.Channels[i], err = createChannelDefinition(v)
		if err != nil {
			return &ChannelDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
		}
	}

	definitions.Message = raw.Message
	definitions.IsSuccess = raw.Message == ""
	return &definitions, nil
}

func createChannelDefinition(raw rawChannelDefinition) (ChannelDefinition, error) {
	d := ChannelDefinition{
		ID:   raw.ID,
		Name: raw.Name,
		Type: raw.Type,
	}
	switch d.Type {
	case slackChannel:
		var slack SlackDetail
		err := json.Unmarshal(raw.Detail, &slack)
		if err != nil {
			return ChannelDefinition{}, errors.Wrapf(err, "failed to unmarshal json")
		}
		d.Detail = slack
	default:
		return ChannelDefinition{}, errors.Errorf("unsupported type: %s", d.Type)
	}
	return d, nil
}

func (c *Channel) createGetRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels", c.UserName),
		Header: map[string]string{userToken: c.Token},
	}, nil
}

type rawChannelDefinitions struct {
	Channels []rawChannelDefinition `json:"channels"`
	Result
}

type rawChannelDefinition struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	Detail json.RawMessage `json:"detail"`
}

// ChannelDefinitions is channel definition list.
type ChannelDefinitions struct {
	Channels []ChannelDefinition `json:"channels"`
	Result
}

// ChannelDefinition is channel definition.
type ChannelDefinition struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Detail interface{} `json:"detail"`
}

// UpdateSlackChannel updates a predefined slack channel.
func (c *Channel) UpdateSlackChannel(channelID, channelName string, detail *SlackDetail) (*Result, error) {
	param, err := c.createUpdateSlackChannelRequestParameter(channelID, channelName, detail)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel update parameter")
	}

	return doRequestAndParseResponse(param)
}

func (c *Channel) createUpdateSlackChannelRequestParameter(id, name string, detail *SlackDetail) (*requestParameter, error) {
	update := &slackChannelCreate{
		Name:   name,
		Type:   slackChannel,
		Detail: detail,
	}
	b, err := json.Marshal(update)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels/%s", c.UserName, id),
		Header: map[string]string{userToken: c.Token},
		Body:   b,
	}, nil
}

// Delete deletes the predefined channel settings.
func (c *Channel) Delete(channelID string) (*Result, error) {
	param, err := c.createDeleteSlackChannelRequestParameter(channelID)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (c *Channel) createDeleteSlackChannelRequestParameter(id string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels/%s", c.UserName, id),
		Header: map[string]string{userToken: c.Token},
		Body:   nil,
	}, nil
}
