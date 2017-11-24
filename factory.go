package runner

import (
	"fmt"
	"strings"
	"tcp-client-runner/abstract"
	"tcp-client-runner/utils/logger"
)

type GameCommander struct {
	tcpClient      abstract.Client
	enableCommands []string
	configStatus   bool
}

func StartGameCommander() GameCommander {
	return GameCommander{
		configStatus: false,
		enableCommands: []string{
			"game", "chat", "login", "help",
		},
	}
}

func (factory *GameCommander) GetClient() abstract.Client {
	return factory.tcpClient
}

func (factory *GameCommander) CreateCommand(input string) abstract.Command {
	if input == "help" {
		return factory.CreateHelpCommand()
	}

	if !factory.configStatus || !factory.tcpClient.IsConnect() {
		return factory.CreateConnectCommand()
	}

	for _, comm := range factory.enableCommands {
		if strings.HasPrefix(input, comm) {
			switch input {
			case "login":
				return factory.CreateLoginCommand()
			case "chat":
				return factory.CreateChatCommand()
			}
		}

	}
	fmt.Println(input)
	return factory.CreateGeneralCommand()
}

func (factory *GameCommander) CreateChatCommand() abstract.Command {
	return &ChatCommand{
		tcpClient: factory.tcpClient,
	}
}
func (factory *GameCommander) CreateGameCommand() abstract.Command {
	return nil
}
func (factory *GameCommander) CreateLoginCommand() abstract.Command {
	return &LoginCommand{
		tcpClient: factory.tcpClient,
	}
}
func (factory *GameCommander) CreateConnectCommand() abstract.Command {
	return &ConnectCommand{
		commander: factory,
	}
}
func (factory *GameCommander) CreateHelpCommand() abstract.Command {
	return &HelpCommand{}
}
func (factory *GameCommander) CreateGeneralCommand() abstract.Command {
	return &GeneralCommand{}
}
func (factory *GameCommander) installClient(tcpClient *TcpClient) {
	factory.tcpClient = tcpClient
	factory.configStatus = true
	go onHandleMessage(tcpClient.message)
}

func onHandleMessage(messageChan chan []byte) {
	for {
		message := <-messageChan
		logger.Info(string(message))
	}
}
