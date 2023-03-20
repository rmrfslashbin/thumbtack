package notes

type NotesCmd struct {
	Byid NotesByIdCmd `cmd:"" help:"Returns a single note."`
	List NotesListCmd `cmd:"" help:"Returns a list of the user's notes."`
}
