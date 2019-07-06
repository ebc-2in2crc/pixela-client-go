package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type user struct {
	UserName string
	Token    string
}

func (u *user) Create(agreeTermsOfService, notMinor bool, thanksCode string) (*Result, error) {
	param, err := u.createCreateRequestParameter(agreeTermsOfService, notMinor, thanksCode)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user create parameter")
	}

	return doRequestAndParseResponse(param)
}

func (u *user) createCreateRequestParameter(agreeTermsOfService, notMinor bool, thanksCode string) (*requestParameter, error) {
	create := &userCreate{
		Token:               u.Token,
		UserName:            u.UserName,
		AgreeTermsOfService: boolToString(agreeTermsOfService),
		NotMinor:            boolToString(notMinor),
		ThanksCode:          thanksCode,
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    APIBaseURL + "/users",
		Header: map[string]string{},
		Body:   b,
	}, nil
}

type userCreate struct {
	Token               string `json:"token"`
	UserName            string `json:"username"`
	AgreeTermsOfService string `json:"AgreeTermsOfService"`
	NotMinor            string `json:"NotMinor"`
	ThanksCode          string `json:"thanksCode,omitempty"`
}

func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func (u *user) Update(newToken, thanksCode string) (*Result, error) {
	param, err := u.createUpdateRequestParameter(newToken, thanksCode)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user update parameter")
	}

	return doRequestAndParseResponse(param)
}

func (u *user) createUpdateRequestParameter(newToken, thanksCode string) (*requestParameter, error) {
	update := userUpdate{
		NewToken:   newToken,
		ThanksCode: thanksCode,
	}
	b, err := json.Marshal(update)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   b,
	}, nil
}

type userUpdate struct {
	NewToken   string `json:"newToken"`
	ThanksCode string `json:"thanksCode,omitempty"`
}

func (u *user) Delete() (*Result, error) {
	param, err := u.createDeleteRequestParameter()
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (u *user) createDeleteRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   []byte{},
	}, nil
}
