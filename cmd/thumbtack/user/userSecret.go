package user

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// UserSecretCmd is the command to get the user's secret RSS key.
type UserSecretCmd struct {
	Json bool `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *UserSecretCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "user secret").
		Str("app_name", ctx.Appname).
		Msg("Running command")

	// Create thumbtack client
	client, err := thumbtack.New(
		thumbtack.WithEndpoint(ctx.Endpoint),
		thumbtack.WithToken(ctx.Token),
		thumbtack.WithLogger(ctx.Log),
		thumbtack.WithUserAgent(ctx.UserAgent),
	)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "user secret").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get the user's secret RSS key
	userSecret, err := client.UserSecret()
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "user secret").
			Str("app_name", ctx.Appname).
			Msg("Failed to get user secret")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(userSecret)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "user secret").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal userSecret")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(userSecret)
	}

	return nil
}
