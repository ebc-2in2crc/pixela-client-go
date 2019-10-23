package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Specify the target to be notified.
const (
	TargetQuantity = "quantity"
)

// Specify the condition used to judge whether to notify or not.
const (
	ConditionGreaterThan = ">"
	ConditionEqual       = "="
	ConditionLessThan    = "<"
	ConditionMultipleOf  = "multipleOf"
)

// A Notification manages communication with the Pixela notification API.
type Notification struct {
	UserName string
	Token    string
	GraphID  string
}

// Create creates a new notification rule.
func (n *Notification) Create(notificationID, notificationName, target, condition, threshold, channelID string) (*Result, error) {
	param, err := n.createCreateNotificationRequestParameter(
		notificationID, notificationName, target, condition, threshold, channelID)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification create parameter")
	}

	return doRequestAndParseResponse(param)
}

func (n *Notification) createCreateNotificationRequestParameter(id, name, target, condition, threshold, channelID string) (*requestParameter, error) {
	create := notificationCreate{
		ID:        id,
		Name:      name,
		Target:    target,
		Condition: condition,
		Threshold: threshold,
		ChannelID: channelID,
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", n.UserName, n.GraphID),
		Header: map[string]string{userToken: n.Token},
		Body:   b,
	}, nil
}

type notificationCreate struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Condition string `json:"condition"`
	Threshold string `json:"threshold"`
	ChannelID string `json:"channelID"`
}

// GetAll get all predefined notifications.
func (n *Notification) GetAll() (*NotificationDefinitions, error) {
	param, err := n.createGetNotificationRequestParameter()
	if err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to create notification get parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions NotificationDefinitions
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

// NotificationDefinitions is notification list.
type NotificationDefinitions struct {
	Notifications []NotificationDefinition `json:"notifications"`
	Result
}

// NotificationDefinition is notification definition.
type NotificationDefinition struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Condition string `json:"condition"`
	Threshold string `json:"threshold"`
	ChannelID string `json:"channelID"`
}

func (n *Notification) createGetNotificationRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", n.UserName, n.GraphID),
		Header: map[string]string{userToken: n.Token},
		Body:   nil,
	}, nil
}

// Update updates predefined notification rule.
func (n *Notification) Update(notificationID, notificationName, target, condition, threshold, channelID string) (*Result, error) {
	param, err := n.createUpdateNotificationRequestParameter(
		notificationID, notificationName, target, condition, threshold, channelID)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification update parameter")
	}

	return doRequestAndParseResponse(param)
}

func (n *Notification) createUpdateNotificationRequestParameter(id, name, target, condition, threshold, channelID string) (*requestParameter, error) {
	create := notificationUpdate{
		Name:      name,
		Target:    target,
		Condition: condition,
		Threshold: threshold,
		ChannelID: channelID,
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", n.UserName, n.GraphID, id),
		Header: map[string]string{userToken: n.Token},
		Body:   b,
	}, nil
}

type notificationUpdate notificationCreate

// Delete deletes predefined notification settings.
func (n *Notification) Delete(notificationID string) (*Result, error) {
	param, err := n.createDeleteNotificationRequestParameter(notificationID)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (n *Notification) createDeleteNotificationRequestParameter(id string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", n.UserName, n.GraphID, id),
		Header: map[string]string{userToken: n.Token},
		Body:   nil,
	}, nil
}
