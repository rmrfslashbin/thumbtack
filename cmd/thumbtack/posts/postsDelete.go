package posts

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsDeleteCmd is the command to delete a bookmark.
type PostsDeleteCmd struct {
	Url  string `name:"url" required:"" help:"URL of bookmark to delete" type:"string"`
	Json bool   `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsDeleteCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts del").
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
			Str("cmd", "posts del").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Delete the bookmark
	del, err := client.PostsDelete(cmd.Url)
	if err != nil {
		if _, ok := err.(*thumbtack.ErrUnexpectedResponse); ok {
			ctx.Log.Error().
				Str("cmd", "posts del").
				Str("app_name", ctx.Appname).
				Str("result", err.(*thumbtack.ErrUnexpectedResponse).ResultCode).
				Msg("Failed to delete bookmark")
		} else {
			ctx.Log.Error().
				Str("cmd", "posts del").
				Str("app_name", ctx.Appname).
				Msg("Failed to delete bookmark")
		}
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(del)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts del").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal delete")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(del)
	}

	return nil
}
