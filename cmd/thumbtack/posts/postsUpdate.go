package posts

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsUpdateCmd is the command Returns the most recent time a bookmark was added, updated or deleted.
type PostsUpdateCmd struct {
	Json bool `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsUpdateCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts update").
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
			Str("cmd", "posts update").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get the last updated bookmark timestamp
	update, err := client.PostsUpdate()
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts update").
			Str("app_name", ctx.Appname).
			Msg("Failed to get bookmarks")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(update)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts update").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal update")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(update)
	}

	return nil
}
