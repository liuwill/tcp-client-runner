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

func startGameCommander() GameCommander {
	// tcpClient = buildTcpClient("")
	// tcpClient.Connect()

	// go startEmitter(tcpClient.message)

	return GameCommander{
		configStatus: false,
		enableCommands: []string{
			"game", "chat", "login",
		},
	}
}

func (factory *GameCommander) CreateCommand(input string) Command {
	if !factory.configStatus {
		return factory.CreateConnectCommand()
	}

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
	return &ConnectCommand{
		commander: factory,
	}
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
