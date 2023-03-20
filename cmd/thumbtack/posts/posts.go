package posts

type PostsCmd struct {
	Add     PostsAddCmd     `cmd:"" help:"Add a bookmark."`
	All     PostsAllCmd     `cmd:"" help:"Get all bookmarks."`
	Dates   PostsDatesCmd   `cmd:"" help:"Get dates with bookmarks."`
	Del     PostsDeleteCmd  `cmd:"" help:"Delete a bookmark."`
	Get     PostsGetCmd     `cmd:"" help:"Get specific bookmarks."`
	Recent  PostsRecentCmd  `cmd:"" help:"Get recent bookmarks."`
	Suggest PostsSuggestCmd `cmd:"" help:"Get suggested tags for a URL."`
	Update  PostsUpdateCmd  `cmd:"" help:"Returns the most recent time a bookmark was added, updated or deleted."`
}
