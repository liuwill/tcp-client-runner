package runner

type Command interface {
	Execute(data map[string]string)
	// Parse(string) map[string]string
	Fields() []string
}

type CommandFactory interface {
	CreateLoginCommand() Command
	CreateChatCommand() Command
	CreateGameCommand() Command
	CreateConnectCommand() Command
	CreateGeneralCommand() Command
}
