package notes

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrfslashbin/thumbtack"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/clictx"
)

// NotesByIdCmd is the command to returns a single note by ID
type NotesByIdCmd struct {
	Id   string `name:"id" required:"" help:"The ID of the note to retrieve" type:"string"`
	Json bool   `name:"json" help:"Output as JSON" default:"false" type:"bool"`
}

// Run runs the command
func (cmd *NotesByIdCmd) Run(ctx *clictx.Context) error {
	// Say hello
	ctx.Log.Debug().
		Str("cmd", "notes byid").
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
			Str("cmd", "notes byid").
			Str("app_name", ctx.Appname).
			Msg("Failed to create client")
		return err
	}

	// Get note with params
	note, err := client.NotesById(cmd.Id)
	if err != nil {
		ctx.Log.Error().
			Str("cmd", "notes byid").
			Str("app_name", ctx.Appname).
			Msg("Failed to get notes byid")
		return err
	}

	if cmd.Json {
		// Print the result as JSON
		data, err := json.Marshal(note)
		if err != nil {
			ctx.Log.Error().
				Str("cmd", "notes byid").
				Str("app_name", ctx.Appname).
				Msg("Failed to marshal note")
			return err
		}
		fmt.Println(string(data))
	} else {
		// Spew the result
		spew.Dump(note)
	}

	return nil
}
