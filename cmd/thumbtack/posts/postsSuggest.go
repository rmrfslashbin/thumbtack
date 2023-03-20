package posts

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsSuggestCmd is the command to return one or more posts on a single day matching the arguments
type PostsSuggestCmd struct {
	Url  string `name:"url" required:"" help:"URL from which to suggest" type:"string"`
	Json bool   `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsSuggestCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts suggest").
		Str("app_name", ctx.Appname).
		Msg("Running command")

	// Create thumbtack client
	client, err := thumbtack.New(
		thumbtack.WithEndpoint(ctx.Endpoint),
		thumbtack.WithToken(*ctx.Token),
		thumbtack.WithLogger(ctx.Log),
		thumbtack.WithUserAgent(*ctx.UserAgent),
	)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts suggest").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get suggestions
	suggests, err := client.PostsSuggest(cmd.Url)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts suggest").
			Str("app_name", ctx.Appname).
			Msg("Failed to get suggestions")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(suggests)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts suggest").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal suggests")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(suggests)
	}
	return nil
}
