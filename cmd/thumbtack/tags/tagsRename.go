package tags

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// TagsRenameCmd is the command to rename a tag
type TagsRenameCmd struct {
	Old  string `name:"old" required:"" help:"Old tag to name" type:"string"`
	New  string `name:"new" required:"" help:"New tag name" type:"string"`
	Json bool   `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *TagsRenameCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "tags rename").
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
			Str("cmd", "tags rename").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Rename a tag
	rename, err := client.TagsRename(&thumbtack.TagsRenameInput{
		Old: &cmd.Old,
		New: &cmd.New,
	})
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "tags rename").
			Str("app_name", ctx.Appname).
			Msg("Failed to rename tag")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(rename)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "tags rename").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal tags")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(rename)
	}

	return nil
}
