package posts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// PostsAddCmd is the command to add a bookmark.
type PostsAddCmd struct {
	Url       string     `name:"url" required:"" help:"URL to bookmark" type:"string"`
	Title     string     `name:"title" required:"" help:"Title of bookmark" type:"string"`
	Descr     *string    `name:"descr" help:"Description of bookmark" type:"string"`
	Replace   bool       `name:"replace" negatable:"" help:"Replace existing bookmark" default:"true" type:"bool"`
	Shared    bool       `name:"shared" negatable:"" help:"Share bookmark with everyone" default:"true" type:"bool"`
	Tags      []string   `name:"tag" help:"Tags to add to bookmark" type:"string"`
	Timestamp *time.Time `name:"timestamp" help:"Timestamp to add bookmark(format: 2006-01-02T15:04:05Z)" type:"date"`
	Json      bool       `name:"json" help:"Output as JSON" default:"false" type:"bool"`
	Unread    *bool      `name:"unread" help:"Mark bookmark as unread" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *PostsAddCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "posts add").
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
			Str("cmd", "posts add").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Add bookmark with params
	add, err := client.PostsAdd(
		&thumbtack.PostsAddInput{
			Url:         &cmd.Url,
			Title:       &cmd.Title,
			Description: cmd.Descr,
			Replace:     &cmd.Replace,
			Shared:      &cmd.Shared,
			Tags:        cmd.Tags,
			Timestamp:   cmd.Timestamp,
			ToRead:      cmd.Unread,
		},
	)
	if err != nil {
		if _, ok := err.(*thumbtack.ErrUnexpectedResponse); ok {
			ctx.Log.Error().
				Str("cmd", "posts add").
				Str("app_name", ctx.Appname).
				Str("result", err.(*thumbtack.ErrUnexpectedResponse).ResultCode).
				Msg("Failed to add bookmark")
		} else {
			ctx.Log.Error().
				Str("cmd", "posts add").
				Str("app_name", ctx.Appname).
				Msg("Failed to add bookmark")
		}
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(add)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "posts add").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal add")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(add)
	}

	return nil
}
