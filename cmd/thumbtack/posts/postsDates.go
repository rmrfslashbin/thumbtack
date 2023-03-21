package posts

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsDatesCmd is the command to return one or more posts on a single day matching the arguments
type PostsDatesCmd struct {
	Tags []string `name:"tag" help:"Get posts with these tags (max: 3)" type:"strings"`
	Json bool     `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsDatesCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts dates").
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
			Str("cmd", "posts dates").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get dates and count of bookmarks
	dates, err := client.PostsDates(cmd.Tags)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts dates").
			Str("app_name", ctx.Appname).
			Msg("Failed to get bookmarks")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(dates)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts dates").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal dates")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(dates)
	}
	return nil
}
