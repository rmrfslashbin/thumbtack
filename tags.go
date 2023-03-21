package thumbtack

import (
	"encoding/json"
	"net/url"
)

// TagsDelete deletes a tag from the user's account
// https://pinboard.in/api/#tags_delete
func (c *Client) TagsDelete(tag string) (*Result, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("tag", tag)

	// Call the endpoint
	tagsDelete, err := c.configs.GetAPI("TagsDelete")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(tagsDelete, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsDelete").
			Str("endpoint", c.endpoint.String()).
			Str("path", tagsDelete).
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

// TagsGet returns a full list of the user's tags along with the number of times they were used.
// https://pinboard.in/api/#tags_get
func (c *Client) TagsGet() (*Tags, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	tagsGet, err := c.configs.GetAPI("TagsGet")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(tagsGet, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsGet").
			Str("endpoint", c.endpoint.String()).
			Str("path", tagsGet).
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
	if input == nil {
		return nil, &ErrInvalidInput{}
	}

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
	tagsRename, err := c.configs.GetAPI("TagsRename")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(tagsRename, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::TagsRename").
			Str("endpoint", c.endpoint.String()).
			Str("path", tagsRename).
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
