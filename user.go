package thumbtack

import (
	"encoding/json"
	"net/url"
)

// UserSecret Returns the user's secret RSS key.
// https://pinboard.in/api/#user_secret
func (c *Client) UserSecret() (*Result, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	userSecret, err := c.configs.GetAPI("UserSecret")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(userSecret, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::UserSecret").
			Str("endpoint", c.endpoint.String()).
			Str("path", userSecret).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	secret := &Result{}
	err = json.Unmarshal(*body, secret)
	if err != nil {
		c.log.Error().Msg("error unmarshalling response")
		return nil, &ErrUnmarshalResponse{
			Body: *body,
			Err:  err,
		}
	}

	return secret, nil
}
