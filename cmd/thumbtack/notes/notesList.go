package notes

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// NotesListCmd is the command to returns a list of the user's notes.
type NotesListCmd struct {
	Json bool `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *NotesListCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "notes list").
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
			Str("cmd", "notes list").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get notes list
	notes, err := client.NotesList()
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "notes list").
			Str("app_name", ctx.Appname).
			Msg("Failed to get notes list")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(notes)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "notes list").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal notes")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(notes)
	}

	return nil
}
