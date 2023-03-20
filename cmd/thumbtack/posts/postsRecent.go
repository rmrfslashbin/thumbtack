package posts

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsRecentCmd is the command to return one or more posts on a single day matching the arguments
type PostsRecentCmd struct {
	Count int      `name:"count" help:"Number of posts to return (max: 100)" default:"15" type:"int"`
	Tags  []string `name:"tag" help:"Get posts with these tags (max: 3)" type:"strings"`
	Json  bool     `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsRecentCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts recent").
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
			Str("cmd", "posts recent").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get bookmarks with params
	bookmarks, err := client.PostsRecent(
		&thumbtack.PostsRecentInput{
			Count: &cmd.Count,
			Tags:  cmd.Tags,
		},
	)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts recent").
			Str("app_name", ctx.Appname).
			Msg("Failed to get bookmarks")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(bookmarks)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts recent").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal bookmarks")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(bookmarks)
	}
	return nil
}
