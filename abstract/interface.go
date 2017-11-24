package abstract

type Command interface {
	Execute(data map[string]string)
	// Parse(string) map[string]string
	Fields() []string
}

type ClientCtrl interface {
	GetClient() Client
	CreateCommand(string) Command
}

type BuilderRegister func() (string, func(clientCtrl ClientCtrl) CommandBuilder)

type CommandBuilder interface {
	Build() Command
	SetClient(client *Client)
}

type CommandFactory interface {
	CreateLoginCommand() Command
	CreateChatCommand() Command
	CreateGameCommand() Command
	CreateHelpCommand() Command
	CreateConnectCommand() Command
	CreateGeneralCommand() Command
}
