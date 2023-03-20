package thumbtack

import (
	"encoding/json"
	"strings"
	"time"
)

// Bookmark is a single Pinboard bookmark
type Bookmark struct {
	/* Example:
	{
	    "href": "https://mastodon.world/@signalapp/109983073760757049",
	    "description": "There is no such thing as a global norm. Infrastructure varies, as do culture and communications expectations. For Signal to be a globally useful alternative to surveillance tech, we need to perform well and offer relevant features, even if this means ver",
	    "extended": "",
	    "meta": "2bb37ba4218de1901281284ffb943c99",
	    "hash": "7231c1174261f287aac11ae26dad3a0a",
	    "time": "2023-03-08T03:58:53Z",
	    "shared": "no",
	    "toread": "yes",
	    "tags": "tag1 tag2 tag3"
	  }
	*/

	// Href is the URL of the bookmark
	Href string `json:"href"`

	// Description is the title of the bookmark
	Description string `json:"description"`

	// Extended is the description of the item. Called 'extended' for backwards compatibility with delicious API
	Extended string `json:"extended"`

	// Meta provides a change detection signature for the bookmark
	Meta string `json:"meta"`

	// Hash is a 32 character hexadecimal MD5 hash of the Href/URL
	Hash string `json:"hash"`

	// Time is the time the bookmark was added
	Time time.Time `json:"time"`

	// Shared is whether the bookmark is public or not
	Shared bool `json:"shared"`

	// ToRead is whether the bookmark is marked as unread
	ToRead bool `json:"toread"`

	// Tags is a list of tags associated with the bookmark
	Tags []string `json:"tags"`
}

// UnmarshalJSON is a custom unmarshaler for the Bookmark struct
func (bookmark *Bookmark) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	for key, value := range data {
		switch key {
		case "href":
			bookmark.Href = value.(string)
		case "description":
			bookmark.Description = value.(string)
		case "extended":
			bookmark.Extended = value.(string)
		case "meta":
			bookmark.Meta = value.(string)
		case "hash":
			bookmark.Hash = value.(string)
		case "time":
			timestamp, err := time.Parse(time.RFC3339, value.(string))
			if err != nil {
				return err
			}
			bookmark.Time = timestamp
		case "shared":
			bookmark.Shared = false
			if value == "yes" {
				bookmark.Shared = true
			}
		case "toread":
			bookmark.ToRead = false
			if value == "yes" {
				bookmark.ToRead = true
			}
		case "tags":
			bookmark.Tags = strings.Split(value.(string), " ")
		}
	}
	return nil
}

// Dates is a list of dates and the number of bookmarks added on that date
type Dates struct {
	/* Example:
	{
		"user": "rmrfslashbin",
		"tag": "",
		"dates": {
			"2023-03-19": 4,
			"2023-03-12": 1,
			"2023-03-10": 3,
			"2023-03-09": 1,
			"2023-03-08": 3,
			"2023-03-07": 1,
			"2023-03-04": 1,
			"2023-03-01": 1
		}
	}
	*/
	User  string         `json:"user"`
	Tag   string         `json:"tag"`
	Dates map[string]int `json:"dates"`
}

// Note is a single Pinboard note
type Note struct {
	Id        string    `json:"id"`
	Hash      string    `json:"hash"`
	Title     string    `json:"title"`
	Length    float64   `json:"length"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UnmarshalJSON is a custom unmarshaler for the Note struct
func (note *Note) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	for key, value := range data {
		switch key {
		case "id":
			note.Id = value.(string)
		case "hash":
			note.Hash = value.(string)
		case "title":
			note.Title = value.(string)
		case "length":
			note.Length = value.(float64)
		case "text":
			note.Text = value.(string)
		case "created_at":
			timestamp, err := time.Parse(time.DateTime, value.(string))
			if err != nil {
				return err
			}
			note.CreatedAt = timestamp
		case "updated_at":
			timestamp, err := time.Parse(time.DateTime, value.(string))
			if err != nil {
				return err
			}
			note.UpdatedAt = timestamp
		}
	}
	return nil
}

// Notes is a list of Pinboard notes
type Notes struct {
	/* Example response:
		{
	  "count": 1,
	  "notes": [
	    {
	      "0": "1e5467e342662e6c239c",
	      "id": "1e5467e342662e6c239c",
	      "1": "e652910a03859fd9e80a",
	      "hash": "e652910a03859fd9e80a",
	      "2": "Test Note 01",
	      "title": "Test Note 01",
	      "3": 40,
	      "length": 40,
	      "4": "2023-03-19 14:35:16",
	      "created_at": "2023-03-19 14:35:16",
	      "5": "2023-03-19 14:35:16",
	      "updated_at": "2023-03-19 14:35:16"
	    }
	  ]
	}
	*/
	Count int    `json:"count"`
	Notes []Note `json:"notes"`
}

// Posts is a list of Pinboard posts
type Posts struct {
	/* Example response:
	{
	  "date": "2023-03-10T01:32:09Z",
	  "user": "rmrfslashbin",
	  "posts": [
	    {
	      "href": "https://social.vivaldi.net/@Dallin/109995001825467835",
	      "description": "I’ve long admired Fred Rogers. I grew up watching “Mister Rogers’ Neighborhood” and consider him a mentor and someone I want to emulate. I’m looking forward to reading his biography “The Good Neighbor: The Life and Work of Fred Rogers” by Ma",
	      "extended": "",
	      "meta": "015246ad973e34e9337dcb134e228d28",
	      "hash": "ca241cedb71816c632a9eae027e92226",
	      "time": "2023-03-10T01:32:09Z",
	      "shared": "no",
	      "toread": "yes",
	      "tags": "books"
	    }
	  ]
	}
	*/
	Date  time.Time  `json:"date"`
	User  string     `json:"user"`
	Posts []Bookmark `json:"posts"`
}

// Result is a general response from the Pinboard API
type Result struct {
	/* Example:
	{"result":"done"}
	{"result":"0417237f06a144c09a5c"}
	{"result_code":"missing url"}
	*/
	Result     string `json:"result,omitempty"`
	ResultCode string `json:"result_code,omitempty"`
}

// Suggestions is a list of suggested tags
type Suggestions struct {
	/* Example response:
		[
	  		{
				"popular": [
					"fonts",
					"css",
					"design"
				]
	  		},
	  		{
	    		"recommended": [
					"typography",
					"web",
					"font",
					"webdesign",
					"via:popular",
					"performance",
					"system"
	    		]
	  		}
		]
	*/
	Popular     []string `json:"popular"`
	Recommended []string `json:"recommended"`
}

// UnmarshalJSON is a custom unmarshaler for the Suggestions struct
func (suggestions *Suggestions) UnmarshalJSON(b []byte) error {
	var data []map[string][]string

	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	for _, value := range data {
		for key, value := range value {
			switch key {
			case "popular":
				suggestions.Popular = value
			case "recommended":
				suggestions.Recommended = value
			}
		}
	}
	return nil
}

// Tags is a list of Pinboard tags
type Tags struct {
	/* Example response:
		{
	 		"books": 1,
	  		"custom": 1,
	  		"haproxy": 2,
	  		"logging": 1,
	  		"stats": 1
		}
	*/
	Tags  map[string]int `json:"tags"`
	Count int            `json:"count"`
}

// UnmarshalJSON is a custom unmarshaler for the Note struct
func (tags *Tags) UnmarshalJSON(b []byte) error {
	var data map[string]int

	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	tags.Tags = make(map[string]int, len(data))
	tags.Count = len(data)
	for key, value := range data {
		tags.Tags[key] = value
	}
	return nil
}

// UpdateTime is the time that the Pinboard account was last updated
type UpdateTime struct {
	/* Example response:
	{"update_time":"2023-03-19T15:57:02Z"}
	*/
	UpdateTime time.Time `json:"update_time"`
}
