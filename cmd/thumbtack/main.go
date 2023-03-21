package main

import (
	"net/url"
	"os"

	"github.com/alecthomas/kong"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/root"
	"github.com/rs/zerolog"
)

const (
	// APP_NAME is the name of the application
	APP_NAME = "thumbtack"
)

// main is the entry point
func main() {
	var err error

	// Set up the logger
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Parse the command line
	var cli root.CLI
	ctx := kong.Parse(&cli)

	// Set up the logger's log level
	// Default to info via the CLI args
	switch cli.LogLevel {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	// Log some start up stuff for debugging
	log.Debug().
		Str("app", APP_NAME).
		Msg("Starting up")

	var endpoint *url.URL
	endpoint = nil
	if cli.Endpoint != nil {
		endpoint, err = url.Parse(*cli.Endpoint)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to parse endpoint")
		}
	}

	config := thumbtack.NewConfig()
	var userAgent string
	if cli.UserAgent != nil {
		userAgent = *cli.UserAgent
	} else {
		userAgent = config.GetUserAgent()
	}

	// Call the Run() method of the selected parsed command.
	err = ctx.Run(
		&clictx.Context{
			Log:       &log,
			Token:     &cli.Token,
			Endpoint:  endpoint,
			Appname:   APP_NAME,
			UserAgent: &userAgent,
		})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run command")
	}
}
