package thumbtack

import (
	"encoding/json"
	"net/url"

	"github.com/rmrfslashbin/thumbtack/internal/constants"
)

// UserSecret Returns the user's secret RSS key.
// https://pinboard.in/api/#user_secret
func (c *Client) UserSecret() (*Result, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	body, err := c.callEndpoint(constants.UserSecret, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::UserSecret").
			Str("endpoint", c.endpoint.String()).
			Str("path", constants.UserSecret).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	secret := &Result{}
	err = json.Unmarshal(*body, secret)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return secret, nil
}
