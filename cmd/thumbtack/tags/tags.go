package tags

type TagsCmd struct {
	All    TagsAllCmd    `cmd:"" help:"Returns all tags."`
	Delete TagsDeleteCmd `cmd:"" help:"Deletes a tag."`
	Rename TagsRenameCmd `cmd:"" help:"Renames a tag."`
}
