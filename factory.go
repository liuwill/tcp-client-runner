package runner

import (
	"fmt"
	"strings"
	"tcp-client-runner/utils/logger"
)

type GameCommander struct {
	tcpClient      *TcpClient
	enableCommands []string
}

func startGameCommander() GameCommander {
	tcpClient := buildTcpClient()
	tcpClient.Connect()

	go startEmitter(tcpClient.message)

	return GameCommander{
		tcpClient: tcpClient,
		enableCommands: []string{
			"game", "chat", "login", "connect",
		},
	}
}

func (factory *GameCommander) CreateCommand(input string) Command {
	for _, comm := range factory.enableCommands {
		if strings.HasPrefix(input, comm) {
			switch input {
			case "login":
				return factory.CreateLoginCommand()
			}
		}

	}
	fmt.Println(input)
	return factory.CreateGeneralCommand()
}

func (factory *GameCommander) CreateChatCommand() Command {
	return nil
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
	return nil
}
func (factory *GameCommander) CreateGeneralCommand() Command {
	return &GeneralCommand{}
}

func startEmitter(messageEmitter chan []byte) {
	for {
		message := <-messageEmitter
		logger.Info(string(message))
	}
}
