package configs

import "strings"

// Configs contains the configuration for the client
type Configs struct {
	// apis is a map of the apis endpoints.
	apis map[string]string

	// endpoint is the root url for the api.
	endpoint string

	// useragent is the user agent string for the client.
	useragent string

	// version is the version of the client.
	version string
}

// ErrApiNotSet is returned when the api is not set
type ErrApiNotSet struct {
	// Msg is the error message
	Msg string

	// Api is the api that was requested
	Api string
}

// Error returns the error message
func (e *ErrApiNotSet) Error() string {
	if e.Msg == "" {
		e.Msg = "requested api is not set or empty"
	}
	if e.Api != "" {
		e.Msg += ": " + e.Api
	}
	return e.Msg
}

// ErrUnknownApi is returned when the api is not known
type ErrUnknownApi struct {
	// Msg is the error message
	Msg string

	// Api is the api that was requested
	Api string
}

// Error returns the error message
func (e *ErrUnknownApi) Error() string {
	if e.Msg == "" {
		e.Msg = "requested api is not known or defined"
	}
	if e.Api != "" {
		e.Msg += ": " + e.Api
	}
	return e.Msg
}

// New returns a new Configs struct
func New() *Configs {
	useragent := "github.com/rmrfslashbin/thumbtack@v"
	version := "1.0.1"

	// Set up the default configs
	return &Configs{
		apis: map[string]string{
			"PostsAdd":     "/posts/add",
			"PostsAll":     "/posts/all",
			"PostsDates":   "/posts/dates",
			"PostsDelete":  "/posts/delete",
			"PostsGet":     "/posts/get",
			"PostsRecent":  "/posts/recent",
			"PostsSuggest": "/posts/suggest",
			"PostsUpdate":  "/posts/update",
			"UserSecret":   "/user/secret",
			"NotesById":    "/notes",
			"NotesList":    "/notes/list",
			"TagsGet":      "/tags/get",
			"TagsDelete":   "/tags/delete",
			"TagsRename":   "/tags/rename",
		},
		endpoint:  "https://api.pinboard.in/v1",
		useragent: useragent + version,
		version:   version,
	}
}

// GetAPI returns the api endpoint or an error if the api is not known.
func (c *Configs) GetAPI(api string) (string, error) {
	if value, ok := c.apis[api]; ok {
		if strings.TrimSpace(value) == "" {
			return "", &ErrApiNotSet{Api: api}
		}
		return value, nil
	} else {
		return "", &ErrUnknownApi{Api: api}
	}
}

// GetUserAgent returns the user agent string.
func (c *Configs) GetUserAgent() string {
	return c.useragent
}

// GetEndpoint returns the endpoint.
func (c *Configs) GetEndpoint() string {
	return c.endpoint
}

// GetVersion returns the version.
func (c *Configs) GetVersion() string {
	return c.version
}

// SetEndpoint sets the endpoint.
func (c *Configs) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

// SetUserAgent sets the user agent string.
func (c *Configs) SetUserAgent(useragent string) {
	c.useragent = useragent
}

// SetVersion sets the version.
func (c *Configs) SetVersion(version string) {
	c.version = version
}

// SetAPI sets the api endpoint or returns an error if the api is not known.
func (c *Configs) SetAPI(api string, value string) error {
	if _, ok := c.apis[api]; !ok {
		return &ErrUnknownApi{Api: api}
	}
	c.apis[api] = value
	return nil
}
