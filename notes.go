package thumbtack

import (
	"encoding/json"
	"net/url"

	"github.com/rmrfslashbin/thumbtack/internal/constants"
)

// NoteById returns a single note
// https://pinboard.in/api/#notes_get
func (c *Client) NotesById(id string) (*Note, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	path := constants.NotesById + "/" + id
	body, err := c.callEndpoint(path, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::NotesById").
			Str("endpoint", c.endpoint.String()).
			Str("path", path).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	note := &Note{}
	err = json.Unmarshal(*body, note)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return note, nil
}

// NotesList returns a list of the user's notes
// https://pinboard.in/api/#notes_list
func (c *Client) NotesList() (*Notes, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	body, err := c.callEndpoint(constants.NotesList, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::NotesList").
			Str("endpoint", c.endpoint.String()).
			Str("path", constants.NotesList).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	notes := &Notes{}
	err = json.Unmarshal(*body, notes)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return notes, nil
}
