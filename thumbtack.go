package thumbtack

// https://pinboard.in/api/

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/rmrfslashbin/thumbtack/internal/configs"
	"github.com/rs/zerolog"
)

// Options for the controller query
type Option func(c *Client)

// Client provides access to the Thumbtack API
type Client struct {
	// configs. the configs for the controller
	configs *configs.Configs

	// dateTimeFormat. the format of the time in the response. dateTimeFormat is always RFC3339
	dateTimeFormat string

	// endpoint. the endpoint is static, but can be overridden for testing
	endpoint *url.URL

	// format. the format of the response. format is always json
	format string

	// logger. if not provided, a default logger will be used
	log *zerolog.Logger

	// token. the token is required for all requests
	token *string

	// userAgent. the userAgent is required for all requests
	userAgent *string
}

// New creates a new Thumbtack client
func New(opts ...Option) (*Client, error) {
	client := &Client{}

	// set up defaults
	client.format = "json"
	client.dateTimeFormat = time.RFC3339 //"2006-01-02T15:04:05Z"

	// apply the list of options to Client
	for _, opt := range opts {
		opt(client)
	}

	// set up logger if not provided
	if client.log == nil {
		log := zerolog.New(os.Stderr).With().Timestamp().Logger()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("No logger provided, using default logger")
		client.log = &log
	}

	if client.configs == nil {
		client.configs = configs.New()
	}

	// set up token if not provided
	if client.token == nil {
		return nil, &ErrNoToken{}
	}

	// set up endpoint
	if client.endpoint == nil {
		endpoint, err := url.Parse(client.configs.GetEndpoint())
		if err != nil {
			return nil, &ErrBadEndpoint{}
		}
		client.endpoint = endpoint
	}

	// set up userAgent if not provided
	if client.userAgent == nil {
		ua := client.configs.GetUserAgent()
		client.userAgent = &ua
	}

	return client, nil
}

func WithConfigs(configs *configs.Configs) Option {
	return func(c *Client) {
		c.configs = configs
	}
}

// WithEndpoint sets the endpoint for the controller
func WithEndpoint(endpoint *url.URL) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithLogger sets the logger for the controller
func WithLogger(log *zerolog.Logger) Option {
	return func(c *Client) {
		c.log = log
	}
}

// WithToken sets the token for the controller
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = &token
	}
}

// WithUserAgent sets the userAgent for the controller
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = &userAgent
	}
}

// callEndpoint calls the endpoint and returns the response body
func (c *Client) callEndpoint(path string, query string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s?%s", c.endpoint.String(), path, query)
	c.log.Debug().
		Str("url", url).
		Msg("calling endpoint")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.log.Error().Msg("failed to create request")
		return nil, err
	}

	req.Header.Add("User-Agent", *c.userAgent)

	res, err := client.Do(req)
	if err != nil {
		c.log.Error().Msg("failed to call endpoint")
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.log.Error().Msg("failed to read response body")
		return nil, err
	}

	// check status code and return error if not 200
	if res.StatusCode != 200 {
		return nil, &ErrBadStatusCode{
			StatusCode: res.StatusCode,
			Status:     res.Status,
		}
	}

	return &body, nil
}
