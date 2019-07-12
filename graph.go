package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Graph manages communication with the Pixela graph API.
type Graph struct {
	UserName string
	Token    string
	GraphID  string
}

// Create creates a new pixelation graph definition.
func (g *Graph) Create(name, unit, quantityType, color, timezone, selfSufficient string, isSecret, publishOptionalData bool) (*Result, error) {
	param, err := g.createCreateRequestParameter(name, unit, quantityType, color, timezone, selfSufficient, isSecret, publishOptionalData)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph create parameter")
	}

	return doRequestAndParseResponse(param)
}

func (g *Graph) createCreateRequestParameter(name, unit, quantityType, color, timezone, selfSufficient string, isSecret, publishOptionalData bool) (*requestParameter, error) {
	create := graphCreate{
		ID:                  g.GraphID,
		Name:                name,
		Unit:                unit,
		Type:                quantityType,
		Color:               color,
		TimeZone:            timezone,
		SelfSufficient:      selfSufficient,
		IsSecret:            isSecret,
		PublishOptionalData: publishOptionalData,
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs", g.UserName),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

type graphCreate struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Unit                string `json:"unit"`
	Type                string `json:"type"`
	Color               string `json:"color"`
	TimeZone            string `json:"timezone"`
	SelfSufficient      string `json:"selfSufficient"`
	IsSecret            bool   `json:"isSecret"`
	PublishOptionalData bool   `json:"publishOptionalData"`
}

// It is the type of quantity to be handled in the graph.
// Only int or float are supported.
const (
	TypeInt   = "int"
	TypeFloat = "float"
)

// Defines the display color of the pixel in the pixelation graph.
// shibafu (green), momiji (red), sora (blue), ichou (yellow), ajisai (purple) and kuro (black) are supported as color kind.
const (
	ColorShibafu = "shibafu"
	ColorMomiji  = "momiji "
	ColorSora    = "sora"
	ColorIchou   = "ichou"
	ColorAjisai  = "ajisai"
	ColorKuro    = "kuro"
)

// If SVG graph with this field increment or decrement is referenced, Pixel of this graph itself will be incremented or decremented.
// It is suitable when you want to record the PVs on a web page or site simultaneously.
// The specification of increment or decrement is the same as Increment a Pixel and Decrement a Pixel with webhook.
// If not specified, it is treated as none .
const (
	SelfSufficientIncrement = "increment"
	SelfSufficientDecrement = "decrement"
	SelfSufficientNone      = "none"
)

// GetAll gets all predefined pixelation graph definitions.
func (g *Graph) GetAll() (*GraphDefinitions, error) {
	param, err := g.createGetAllRequestParameter()
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to create get all graph parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions GraphDefinitions
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

func (g *Graph) createGetAllRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs", g.UserName),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// GraphDefinitions is graph definition list.
type GraphDefinitions struct {
	Graphs []GraphDefinition `json:"graphs"`
	Result
}

// GraphDefinition is graph definition.
type GraphDefinition struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Unit                string   `json:"unit"`
	Type                string   `json:"type"`
	Color               string   `json:"color"`
	TimeZone            string   `json:"timezone"`
	PurgeCacheURLs      []string `json:"purgeCacheURLs"`
	SelfSufficient      string   `json:"selfSufficient"`
	IsSecret            bool     `json:"isSecret"`
	PublishOptionalData bool     `json:"publishOptionalData"`
}

// GetSVG get a graph expressed in SVG format diagram that based on the registered information.
func (g *Graph) GetSVG(date, mode string) (string, error) {
	param, err := g.createGetSVGRequestParameter(date, mode)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create get svg parameter")
	}

	b, err := mustDoRequest(param)
	if err != nil {
		return "", errors.Wrapf(err, "failed to do request")
	}

	return string(b), nil
}

func (g *Graph) createGetSVGRequestParameter(date, mode string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s?date=%s&mode=%s", g.UserName, g.GraphID, date, mode),
		Header: map[string]string{userToken: token},
		Body:   []byte{},
	}, nil
}

// Specify the graph display mode.
// Supported modes are short (for displaying only about 90 days) and line .
const (
	ModeShort = "short"
	ModeLine  = "line"
)

// URL displays the details of the graph in html format.
func (g *Graph) URL() string {
	return fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s.html", g.UserName, g.GraphID)
}

// GraphsURL displays graph list by detail in html format.
func (g *Graph) GraphsURL() string {
	return fmt.Sprintf(APIBaseURL+"/users/%s/graphs.html", g.UserName)
}

// Stats is various statistics based on the registered information.
type Stats struct {
	TotalPixelsCount int     `json:"totalPixelsCount"`
	MaxQuantity      int     `json:"maxQuantity"`
	MinQuantity      int     `json:"minQuantity"`
	TotalQuantity    int     `json:"totalQuantity"`
	AvgQuantity      float64 `json:"avgQuantity"`
	TodaysQuantity   int     `json:"todaysQuantity"`
	Result
}

// Stats gets various statistics based on the registered information.
func (g *Graph) Stats() (*Stats, error) {
	param, err := g.createStatsRequestParameter()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create graph stats request parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to do request")
	}

	var stats Stats
	if err := json.Unmarshal(b, &stats); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal json")
	}

	stats.IsSuccess = stats.Message == ""
	return &stats, nil
}

func (g *Graph) createStatsRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/stats", g.UserName, g.GraphID),
		Header: map[string]string{},
		Body:   []byte{},
	}, nil
}

// Update updates predefined pixelation graph definitions.
// The items that can be updated are limited as compared with the pixelation graph definition creation.
func (g *Graph) Update(name, unit, color, timezone string, purgeCacheUrls []string, selfSufficient string, isSecret bool, publishOptionalData bool) (*Result, error) {
	param, err := g.createUpdateRequestParameter(name, unit, color, timezone, purgeCacheUrls, selfSufficient, isSecret, publishOptionalData)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph update parameter")
	}

	return doRequestAndParseResponse(param)
}

func (g *Graph) createUpdateRequestParameter(name, unit, color, timezone string, purgeCacheUrls []string, selfSufficient string, isSecret, publishOptionalData bool) (*requestParameter, error) {
	update := graphUpdate{
		Name:                name,
		Unit:                unit,
		Color:               color,
		TimeZone:            timezone,
		PurgeCacheURLs:      purgeCacheUrls,
		SelfSufficient:      selfSufficient,
		IsSecret:            isSecret,
		PublishOptionalData: publishOptionalData,
	}
	b, err := json.Marshal(update)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", g.UserName, g.GraphID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

type graphUpdate struct {
	Name                string   `json:"name"`
	Unit                string   `json:"unit"`
	Color               string   `json:"color"`
	TimeZone            string   `json:"timezone"`
	PurgeCacheURLs      []string `json:"purgeCacheURLs"`
	SelfSufficient      string   `json:"selfSufficient"`
	IsSecret            bool     `json:"isSecret"`
	PublishOptionalData bool     `json:"publishOptionalData"`
}

// Delete deletes the predefined pixelation graph definition.
func (g *Graph) Delete() (*Result, error) {
	param, err := g.createDeleteRequestParameter()
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (g *Graph) createDeleteRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", g.UserName, g.GraphID),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// GetPixelDates gets a Date list of Pixel registered in the graph specified by graphID.
// You can specify a period with from and to parameters.
//
// If you do not specify both from and to;
// You will get a list of 365 days ago from today.
//
// If you specify from only;
// You will get a list of 365 days from from date.
//
// If you specify to only;
// You will get a list of 365 days ago from to date.
//
// If you specify both from andto;
// You will get a list you specify.
// You can not specify a period greater than 365 days.
func (g *Graph) GetPixelDates(from, to string) (*Pixels, error) {
	param, err := g.createGetPixelDatesRequestParameter(from, to)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to create get pixel dates parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to do request")
	}

	var pixels Pixels
	if err := json.Unmarshal(b, &pixels); err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	pixels.IsSuccess = pixels.Message == ""
	return &pixels, nil
}

// Pixels is Date list of Pixel registered in the graph.
type Pixels struct {
	Pixels []string `json:"pixels"`
	Result
}

func (g *Graph) createGetPixelDatesRequestParameter(from, to string) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/pixels?from=%s&to=%s", g.UserName, g.GraphID, from, to),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}
