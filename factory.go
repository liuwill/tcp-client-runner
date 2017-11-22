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

// ** 创建游戏命令 ** //

// ** 执行游戏命令 ** //
type Invoker struct {
	command Command
}

func (invoker *Invoker) SetCommand(command Command) {
	invoker.command = command
}

func (invoker *Invoker) Action(input string) {
	invokerData := invoker.Parse(input, invoker.command)
	invoker.command.Execute(invokerData)
}
func (invoker *Invoker) Parse(input string, command Command) map[string]string {
	input = strings.Replace(input, "  ", " ", -1)
	inputPiece := strings.Split(input, " ")
	result := make(map[string]string)

	for index, field := range command.Fields() {
		if index < len(inputPiece) && len(inputPiece[index]) > 0 {
			result[field] = inputPiece[index]
		}
	}

	if len(command.Fields()) == 0 {
		result["content"] = strings.Join(inputPiece[1:], " ")
		return result
	}

	return result
}

func startEmitter(messageEmitter chan []byte) {
	for {
		message := <-messageEmitter
		logger.Info(string(message))
	}
}
