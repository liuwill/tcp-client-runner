package runner

import (
	"fmt"
	"strings"
	"tcp-client-runner/utils/logger"
)

type GameCommander struct {
	tcpClient      *TcpClient
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

func (factory *GameCommander) GetClient() *TcpClient {
	return factory.tcpClient
}

func (factory *GameCommander) CreateCommand(input string) Command {
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

func (factory *GameCommander) CreateChatCommand() Command {
	return &ChatCommand{
		tcpClient: factory.tcpClient,
	}
}
func (factory *GameCommander) CreateGameCommand() Command {
	return nil
}
func (factory *GameCommander) CreateLoginCommand() Command {
	return &LoginCommand{
		tcpClient: factory.tcpClient,
	}
}
func (factory *GameCommander) CreateConnectCommand() Command {
	return &ConnectCommand{
		commander: factory,
	}
}
func (factory *GameCommander) CreateHelpCommand() Command {
	return &HelpCommand{}
}
func (factory *GameCommander) CreateGeneralCommand() Command {
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
