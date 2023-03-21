package thumbtack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Functions relate to the /post API endpoint

// PostsAddInput is the input for the PostsAdd function
type PostsAddInput struct {
	// Url is the URL of the item.
	// Required.
	Url *string

	// Title is the title of the item.
	// Required.
	Title *string

	// Description is the description of the item.
	// backwards compatibility with delicious API.
	Description *string

	// Replace any existing bookmark with this URL.
	// Default is true.
	// If set to false, will throw an error if bookmark exists.
	Replace *bool

	// Shared is a boolean indicating whether the bookmark is public.
	// Default is true unless user has enabled the "save all bookmarks
	// as private" user setting, in which case default is false.
	Shared *bool

	// Tags is a list of up to 100 tags.
	Tags []string

	// Timestamp is the creation time for this bookmark.
	// Defaults to current time.
	// Datestamps more than 10 minutes ahead of server time will be reset to current server time.
	Timestamp *time.Time

	// ToRead marks the bookmark as unread.
	// Default is false
	ToRead *bool
}

// PostsAdd Add a bookmark
// https://pinboard.in/api/#posts_add
func (c *Client) PostsAdd(input *PostsAddInput) (*Result, error) {
	// Input validation
	if input == nil {
		return nil, &ErrInvalidInput{}
	}

	if input.Url == nil {
		return nil, &ErrMissingInputField{Field: "Url"}
	}

	if input.Title == nil {
		return nil, &ErrMissingInputField{Field: "Title"}
	}

	if len(input.Tags) > 100 {
		return nil, &ErrInvalidInput{Msg: "tags must be less than or equal to 100"}
	}

	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("url", *input.Url)
	v.Set("description", *input.Title) // description is the title

	if input.Description != nil {
		v.Set("extended", *input.Description) // extended is the description
	}

	if input.Replace != nil {
		replace := "yes"
		if !*input.Replace {
			replace = "no"
		}
		v.Set("replace", replace)
	}

	if input.Shared != nil {
		shared := "yes"
		if !*input.Shared {
			shared = "no"
		}
		v.Set("shared", shared)
	}

	if len(input.Tags) > 0 {
		v.Set("tags", strings.Join(input.Tags, " "))
	}

	if input.Timestamp != nil {
		if input.Timestamp.After(time.Now().Add(10 * time.Minute)) {
			return nil, &ErrInvalidInput{Msg: "timestamp must be less than or equal to 10 minutes ahead of server time"}
		}
		v.Set("dt", input.Timestamp.Format(time.RFC3339))
	}

	if input.ToRead != nil {
		toRead := "no"
		if *input.ToRead {
			toRead = "yes"
		}
		v.Set("toread", toRead)
	}

	// Call the endpoint
	postsAdd, err := c.configs.GetAPI("PostsAdd")
	if err != nil {
		return nil, err
	}

	body, err := c.callEndpoint(postsAdd, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsAdd").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsAdd).
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

	if result.ResultCode != "done" {
		return nil, &ErrUnexpectedResponse{ResultCode: result.ResultCode}
	}

	return result, nil
}

// PostsAllInput is the input for the PostsAll function
type PostsAllInput struct {
	// fromdt	datetime	return only bookmarks created after this time
	FromDT *time.Time `json:"fromdt"`

	// meta	int	include a change detection signature for each bookmark
	Meta *bool `json:"meta"`

	// results	int	number of results to return. Default is all
	Results *int `json:"results"`

	// start	int	offset value (default is 0)
	Start *int `json:"start"`

	// tag	tag	filter by up to three tags
	Tags []string `json:"tags"`

	// todt	datetime	return only bookmarks created before this time
	ToDT *time.Time `json:"todt"`
}

// PostsAll Returns all bookmarks in the user's account.
// https://pinboard.in/api/#posts_all
func (c *Client) PostsAll(input *PostsAllInput) (*[]Bookmark, error) {
	if input == nil {
		input = &PostsAllInput{}
	}

	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	if input.FromDT != nil {
		v.Set("fromdt", input.FromDT.Format(c.dateTimeFormat))
	}

	// Convert the bool to a string
	meta := "yes"
	if input.Meta != nil && !*input.Meta {
		meta = "no"
	}
	v.Set("meta", meta)

	if input.Results != nil {
		v.Set("results", fmt.Sprint(*input.Results))
	}

	if input.Start != nil {
		v.Set("start", fmt.Sprint(*input.Start))
	}

	// Convert string slice to a string; check for max length

	if len(input.Tags) > 3 {
		return nil, &ErrInvalidInput{Msg: "tags must be less than or equal to 3"}
	}
	if len(input.Tags) > 0 {
		v.Set("tag", strings.Join(input.Tags, " "))
	}

	if input.ToDT != nil {
		v.Set("todt", input.ToDT.Format(c.dateTimeFormat))
	}

	// Call the endpoint
	postsAll, err := c.configs.GetAPI("PostsAll")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsAll, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsAll").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsAll).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	bookmarks := &[]Bookmark{}
	err = json.Unmarshal(*body, bookmarks)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return bookmarks, nil
}

// PostsDates returns a list of dates with the number of posts at each date.
// https://pinboard.in/api/#posts_dates
func (c *Client) PostsDates(tags []string) (*Dates, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Convert string slice to a string; check for max length
	if len(tags) > 3 {
		return nil, &ErrInvalidInput{Msg: "tags must be less than or equal to 3"}
	}
	if len(tags) > 0 {
		v.Set("tag", strings.Join(tags, " "))
	}

	// Call the endpoint
	postsDates, err := c.configs.GetAPI("PostsDates")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsDates, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsDates").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsDates).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	dates := &Dates{}
	err = json.Unmarshal(*body, dates)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return dates, nil
}

// PostsDelete deletes a bookmark
// https://pinboard.in/api/#posts_delete
func (c *Client) PostsDelete(urlToDelete string) (*Result, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("url", urlToDelete)

	// Call the endpoint
	postsDelete, err := c.configs.GetAPI("PostsDelete")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsDelete, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsDelete").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsDelete).
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

	if result.ResultCode != "done" {
		return nil, &ErrUnexpectedResponse{ResultCode: result.ResultCode}
	}

	return result, nil
}

// PostsGetInput is the input for the PostsGet function
type PostsGetInput struct {
	// Date return results bookmarked on this day
	Date *time.Time

	// Meta	includes a change detection signature in a meta attribute
	// Default is true
	Meta *bool

	// Tags filter by up to three tags
	Tags []string

	// url return only this bookmark
	URL *string
}

// PostsGet returns one or more posts on a single day matching the arguments.
// If no date or url is given, date of most recent bookmark will be used.
// https://pinboard.in/api/#posts_get
func (c *Client) PostsGet(input *PostsGetInput) (*Posts, error) {
	if input == nil {
		input = &PostsGetInput{}
	}

	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Convert the bool to a string
	meta := "yes"
	if input.Meta != nil && !*input.Meta {
		meta = "no"
	}
	v.Set("meta", meta)

	// Convert string slice to a string; check for max length
	if len(input.Tags) > 3 {
		return nil, &ErrInvalidInput{Msg: "tags must be less than or equal to 3"}
	}
	if len(input.Tags) > 0 {
		v.Set("tag", strings.Join(input.Tags, " "))
	}

	if input.Date != nil {
		v.Set("dt", input.Date.Format(time.DateOnly))
	}

	if input.URL != nil {
		v.Set("url", *input.URL)
	}

	// Call the endpoint
	postsGet, err := c.configs.GetAPI("PostsGet")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsGet, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsGet").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsGet).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	bookmarks := &Posts{}
	err = json.Unmarshal(*body, bookmarks)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return bookmarks, nil
}

// PostsRecentInput is the input for the PostsRecent function
type PostsRecentInput struct {
	// Count number of posts to return, default is 15, max is 100
	Count *int

	// Tags filter by up to three tags
	Tags []string
}

// PostsRecent returns recent posts, filtered by tag.
// https://pinboard.in/api/#posts_recent
func (c *Client) PostsRecent(input *PostsRecentInput) (*Posts, error) {
	if input == nil {
		input = &PostsRecentInput{}
	}

	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	count := 15
	if input.Count != nil {
		count = *input.Count
	}
	if count > 100 || count < 1 {
		return nil, &ErrInvalidInput{Msg: "count must be greater then 1 and less than or equal to 100"}
	}
	v.Set("count", strconv.Itoa(count))

	// Convert string slice to a string; check for max length
	if len(input.Tags) > 3 {
		return nil, &ErrInvalidInput{Msg: "tags must be less than or equal to 3"}
	}
	if len(input.Tags) > 0 {
		v.Set("tag", strings.Join(input.Tags, " "))
	}

	// Call the endpoint
	postsRecent, err := c.configs.GetAPI("PostsRecent")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsRecent, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsRecent").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsRecent).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	bookmarks := &Posts{}
	err = json.Unmarshal(*body, bookmarks)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return bookmarks, nil
}

// PostsSuggest returns a list of popular tags and recommended tags for a given URL.
// https://pinboard.in/api/#posts_suggest
func (c *Client) PostsSuggest(urlToSuggest string) (*Suggestions, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)
	v.Set("url", urlToSuggest)

	// Call the endpoint
	postsSuggest, err := c.configs.GetAPI("PostsSuggest")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsSuggest, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsSuggest").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsSuggest).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	suggestions := &Suggestions{}
	err = json.Unmarshal(*body, suggestions)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return suggestions, nil
}

// PostsUpdate Returns the most recent time a bookmark was added, updated or deleted.
// Use this before calling posts/all to see if the data has changed since the last fetch.
// https://pinboard.in/api/#posts_update
func (c *Client) PostsUpdate() (*UpdateTime, error) {
	// Set up the query parameters
	v := url.Values{}
	v.Set("format", c.format)
	v.Set("auth_token", *c.token)

	// Call the endpoint
	postsUpdate, err := c.configs.GetAPI("PostsUpdate")
	if err != nil {
		return nil, err
	}
	body, err := c.callEndpoint(postsUpdate, v.Encode())
	if err != nil {
		c.log.Error().
			Str("function", "thumbtack::PostsUpdate").
			Str("endpoint", c.endpoint.String()).
			Str("path", postsUpdate).
			Str("query", v.Encode()).
			Msg("error calling endpoint")
		return nil, err
	}

	// Unmarshal the response
	updateTime := &UpdateTime{}
	err = json.Unmarshal(*body, updateTime)
	if err != nil {
		c.log.Error().Msg("error unmashalling response")
		return nil, err
	}

	return updateTime, nil
}
