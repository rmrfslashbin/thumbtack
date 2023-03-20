package root

import (
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/notes"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/posts"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/tags"
	"github.com/rmrfslashbin/thumbtack/cmd/thumbtack/user"
)

// CLI is the command line interface
type CLI struct {
	// Global flags/args
	LogLevel  string  `name:"loglevel" env:"LOGLEVEL" default:"info" enum:"panic,fatal,error,warn,info,debug,trace" help:"Set the log level."`
	Endpoint  *string `name:"endpoint" env:"ENDPOINT" help:"Set the API endpoint."`
	Token     string  `name:"token" env:"TOKEN" required:"" help:"Set the API token."`
	UserAgent *string `name:"useragent" env:"USERAGENT" help:"Set the User-Agent header."`

	// Commands
	Notes notes.NotesCmd `cmd:"" help:"Notes commands."`
	Posts posts.PostsCmd `cmd:"" help:"Posts commands."`
	Tags  tags.TagsCmd   `cmd:"" help:"Tags commands."`
	User  user.UserCmd   `cmd:"" help:"User commands."`
}
