package thumbtack

import (
	"encoding/json"
	"net/url"
)

// NoteById returns a single note
// https://pinboard.in/api/#notes_get
func (c *Client) NotesById(id string) (*Note, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	notesById, err := c.configs.GetAPI("NotesById")
	if err != nil {
		return nil, err
	}
	path := notesById + "/" + id
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
		c.log.Error().Msg("error unmarshalling response")
		return nil, &ErrUnmarshalResponse{
			Body: *body,
			Err:  err,
		}
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
	notesList, err := c.configs.GetAPI("NotesList")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(notesList, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::NotesList").
			Str("endpoint", c.endpoint.String()).
			Str("path", notesList).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	notes := &Notes{}
	err = json.Unmarshal(*body, notes)
	if err != nil {
		c.log.Error().Msg("error unmarshalling response")
		return nil, &ErrUnmarshalResponse{
			Body: *body,
			Err:  err,
		}
	}

	return notes, nil
}
