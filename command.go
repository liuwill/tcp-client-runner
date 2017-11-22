package runner

import (
	"encoding/json"
	"fmt"
	"tcp-client-runner/utils/logger"
)

type Command interface {
	Execute(data map[string]string)
	// Parse(string) map[string]string
	Fields() []string
}

type ChatCommand struct {
	tcpClient *TcpClient
}
type GeneralCommand struct{}

func (command *GeneralCommand) Execute(data map[string]string) {
	bytes, _ := json.Marshal(data)
	logger.Info(string(bytes))
}
func (command *GeneralCommand) Fields() []string {
	return []string{
		"command", "content",
	}
}

type LoginCommand struct {
	tcpClient *TcpClient
}

func (command *LoginCommand) Execute(data map[string]string) {
	if command.tcpClient.IsLogin() {
		logger.Warning("You have Login now!")
		return
	}

	fmt.Println("Please Set Username For Login!")
	username := ReadLine("Enter Username", "visitor")
	protocol := ReadLine("Enter Protocol", "json")
	logger.Warning(username + ":" + protocol)
	// bytes, _ := json.Marshal(data)
	// logger.Info(string(bytes))
}
func (command *LoginCommand) Fields() []string {
	return []string{}
}

type ConnectCommand struct {
	tcpClient *TcpClient
}

type CommandFactory interface {
	CreateLoginCommand() Command
	CreateChatCommand() Command
	CreateGameCommand() Command
	CreateConnectCommand() Command
	CreateGeneralCommand() Command
}

