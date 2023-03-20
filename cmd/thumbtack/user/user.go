package user

type UserCmd struct {
	Secret UserSecretCmd `cmd:"" help:"Get the user's secret RSS key."`
}
