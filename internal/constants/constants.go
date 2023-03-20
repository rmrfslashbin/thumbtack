package constants

const (
	// Package constants
	Useragent = "github.com/rmrfslashbin/thumbtack@v" + Version
	Version   = "1.0.1"

	// Default supported endpoint
	Endpoint = "https://api.pinboard.in/v1"

	// Posts endpoints
	PostsAdd     = "/posts/add"
	PostsAll     = "/posts/all"
	PostsDates   = "/posts/dates"
	PostsDelete  = "/posts/delete"
	PostsGet     = "/posts/get"
	PostsRecent  = "/posts/recent"
	PostsSuggest = "/posts/suggest"
	PostsUpdate  = "/posts/update"

	// User endpoints
	UserSecret = "/user/secret"

	// Notes endpoints
	NotesById = "/notes"
	NotesList = "/notes/list"

	// Tags endpoints
	TagsGet    = "/tags/get"
	TagsDelete = "/tags/delete"
	TagsRename = "/tags/rename"
)
