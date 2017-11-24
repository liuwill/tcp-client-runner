package runner

type Command interface {
	Execute(data map[string]string)
	// Parse(string) map[string]string
	Fields() []string
}

type ClientCtrl interface {
	GetClient() *TcpClient
	CreateCommand(string) Command
}

type CommandBuilder interface {
	Build() Command
	SetClient(client *TcpClient)
}

type CommandFactory interface {
	CreateLoginCommand() Command
	CreateChatCommand() Command
	CreateGameCommand() Command
	CreateHelpCommand() Command
	CreateConnectCommand() Command
	CreateGeneralCommand() Command
}
