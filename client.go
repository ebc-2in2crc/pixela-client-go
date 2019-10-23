package pixela

// A Client manages communication with the Pixela User API.
type Client struct {
	UserName string
	Token    string
}

// NewClient return a new Client instance.
func NewClient(userName, token string) *Client {
	return &Client{UserName: userName, Token: token}
}

// CreateUser creates a new Pixela user.
func (c *Client) CreateUser(agreeTermsOfService, notMinor bool, thanksCode string) (*Result, error) {
	return c.user().Create(agreeTermsOfService, notMinor, thanksCode)
}

func (c *Client) user() *user {
	return &user{UserName: c.UserName, Token: c.Token}
}

// UpdateUser updates the authentication token for the specified user.
func (c *Client) UpdateUser(newToken, thanksCode string) (*Result, error) {
	result, err := c.user().Update(newToken, thanksCode)
	if err == nil && result.IsSuccess {
		c.Token = newToken
	}
	return result, err
}

// DeleteUser deletes the specified registered user.
func (c *Client) DeleteUser() (*Result, error) {
	return c.user().Delete()
}

// Channel returns a new Pixela channel API client.
func (c *Client) Channel() *Channel {
	return &Channel{UserName: c.UserName, Token: c.Token}
}

// Graph returns a new Pixela graph API client.
func (c *Client) Graph(graphID string) *Graph {
	return &Graph{UserName: c.UserName, Token: c.Token, GraphID: graphID}
}

// Pixel returns a new Pixela pixel API client.
func (c *Client) Pixel(graphID string) *Pixel {
	return &Pixel{UserName: c.UserName, Token: c.Token, GraphID: graphID}
}

// Notification returns a new Pixela notification API client.
func (c *Client) Notification(graphID string) *Notification {
	return &Notification{UserName: c.UserName, Token: c.Token, GraphID: graphID}
}

// Webhook returns a new Pixela webhook API client.
func (c *Client) Webhook() *Webhook {
	return &Webhook{UserName: c.UserName, Token: c.Token}
}
