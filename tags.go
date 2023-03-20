package thumbtack

import (
	"encoding/json"
	"net/url"

	"github.com/rmrfslashbin/thumbtack/internal/constants"
)

// NotesList returns a list of the user's notes
// https://pinboard.in/api/#notes_list
func (c *Client) TagsGet() (*Tags, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	body, err := c.callEndpoint(constants.TagsGet, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsGet").
			Str("endpoint", c.endpoint.String()).
			Str("path", constants.TagsGet).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	tags := &Tags{}
	err = json.Unmarshal(*body, tags)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return tags, nil
}

// TagsDelete deletes a tag from the user's account
// https://pinboard.in/api/#tags_delete
func (c *Client) TagsDelete(tag string) (*Result, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("tag", tag)

	// Call the endpoint
	body, err := c.callEndpoint(constants.TagsDelete, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsDelete").
			Str("endpoint", c.endpoint.String()).
			Str("path", constants.TagsDelete).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	result := &Result{}
	err = json.Unmarshal(*body, result)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	if result.Result != "done" {
		return nil, &ErrUnexpectedResponse{ResultCode: result.Result}
	}

	return result, nil
}

// TagsRenameInput is the input for the TagsRename function
type TagsRenameInput struct {
	// Old is the tag to rename
	Old *string `json:"old"`

	// New is the new name for the tag
	New *string `json:"new"`
}

// TagsRename renames a tag
// https://pinboard.in/api/#tags_rename
func (c *Client) TagsRename(input *TagsRenameInput) (*Result, error) {
	if input.Old == nil {
		return nil, &ErrMissingInputField{Field: "old"}
	}

	if input.New == nil {
		return nil, &ErrMissingInputField{Field: "new"}
	}

	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("old", *input.Old)
	v.Set("new", *input.New)

	// Call the endpoint
	body, err := c.callEndpoint(constants.TagsRename, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsRename").
			Str("endpoint", c.endpoint.String()).
			Str("path", constants.TagsRename).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	result := &Result{}
	err = json.Unmarshal(*body, result)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	if result.Result != "done" {
		return nil, &ErrUnexpectedResponse{ResultCode: result.Result}
	}

	return result, nil
}
