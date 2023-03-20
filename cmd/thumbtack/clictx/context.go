package clictx

import (
	"net/url"

	"github.com/rs/zerolog"
)

// Context is used to pass context/global configs to the commands
type Context struct {
	// Appname is the name of the application
	Appname string

	// Endpoint is the endpoint to use
	Endpoint *url.URL

	// log is the logger
	Log *zerolog.Logger

	// Token is the token to use
	Token *string

	// UserAgent is the user agent to use
	UserAgent *string
}
