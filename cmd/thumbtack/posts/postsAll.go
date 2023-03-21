package posts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsAllCmd is the command to get all posts
type PostsAllCmd struct {
	From    *time.Time `name:"from" help:"Get posts from this date/time (format: 2006-01-02T15:04:05Z)" type:"date"`
	Meta    bool       `name:"meta" help:"Get meta data for posts" default:"false" type:"bool"`
	Results *int       `name:"results" help:"Get this many results" type:"int"`
	Start   *int       `name:"start" help:"Start at this result" type:"int"`
	Tags    []string   `name:"tags" help:"Get posts with these tags" type:"strings"`
	To      *time.Time `name:"to" help:"Get posts to this date/time (format: 2006-01-02T15:04:05Z)" type:"date"`
	Json    bool       `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsAllCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts all").
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
			Str("cmd", "posts all").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get bookmarks with params
	bookmarks, err := client.PostsAll(
		&thumbtack.PostsAllInput{
			FromDT:  cmd.From,
			ToDT:    cmd.To,
			Meta:    &cmd.Meta,
			Results: cmd.Results,
			Start:   cmd.Start,
			Tags:    cmd.Tags,
		},
	)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "posts all").
			Str("app_name", ctx.Appname).
			Msg("Failed to get bookmarks")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(bookmarks)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts all").
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
