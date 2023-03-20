package tags

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// TagsDeleteCmd is the command to delete a tag
type TagsDeleteCmd struct {
	Tag  string `name:"tag" required:"" help:"Tag to delete" type:"string"`
	Json bool   `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *TagsDeleteCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "tags delete").
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
			Str("cmd", "tags delete").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Delete a tag
	delete, err := client.TagsDelete(cmd.Tag)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "tags delete").
			Str("app_name", ctx.Appname).
			Msg("Failed to delete tag")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(delete)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "tags delete").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal tags")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(delete)
	}

	return nil
}
