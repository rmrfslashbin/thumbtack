package tags

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// TagsAllCmd is the command to get all tags
type TagsAllCmd struct {
	Json bool `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *TagsAllCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "tags all").
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
			Str("cmd", "tags all").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get all tags
	tags, err := client.TagsGet()
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "tags all").
			Str("app_name", ctx.Appname).
			Msg("Failed to get tags")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(tags)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "tags all").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal tags")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(tags)
	}

	return nil
}
