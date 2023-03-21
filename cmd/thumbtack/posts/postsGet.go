package posts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsGetCmd is the command to return one or more posts on a single day matching the arguments
type PostsGetCmd struct {
	Date *string  `name:"date" help:"Get posts from this date (format: 2006-01-02)" type:"string"`
	Meta bool     `name:"meta" negatable:"" help:"Get meta data for posts" default:"true" type:"bool"`
	Tags []string `name:"tag" help:"Get posts with these tags (max: 3)" type:"strings"`
	Url  *string  `name:"url" help:"Get posts with this URL" type:"string"`
	Json bool     `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsGetCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts get").
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
			Str("cmd", "posts get").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	var date time.Time
	if cmd.Date != nil {
		date, err = time.Parse(time.DateOnly, *cmd.Date)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts get").
				Str("app_name", ctx.Appname).
				Msg("error parsing date")
			return err
		}
	}
	// Get bookmarks with params
	bookmarks, err := client.PostsGet(
		&thumbtack.PostsGetInput{
			Date: &date,
			Meta: &cmd.Meta,
			Tags: cmd.Tags,
			URL:  cmd.Url,
		},
	)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts get").
			Str("app_name", ctx.Appname).
			Msg("Failed to get bookmarks")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(bookmarks)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts get").
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
