package runner

import (
	"fmt"
	"tcp-client-runner/abstract"
	"tcp-client-runner/utils/io"
	"tcp-client-runner/utils/logger"
)

type EchoCommand struct {
	tcpClient abstract.Client
}

func (command *EchoCommand) Execute(data map[string]string) {
	if command.tcpClient == nil || !command.tcpClient.IsConnect() {
		logger.Error("You need to connect with server")
		return
	}

	fmt.Println("# Echo Client Startup! ('exit' or 'quit' to leave)")
	for {
		input := io.ReadLine("Echo > ", "")
		if len(input) <= 0 {
			continue
		} else if input == "exit" || input == "quit" {
			break
		}

		command.tcpClient.SendBytes([]byte(input))
	}
}

func (command *EchoCommand) Fields() []string {
	return []string{}
}

type EchoCommandBuilder struct {
	clientCtrl abstract.ClientCtrl
}

func (builder *EchoCommandBuilder) Build() abstract.Command {
	return &EchoCommand{
		tcpClient: builder.clientCtrl.GetClient(),
	}
}

func (builder *EchoCommandBuilder) SetClientCtrl(clientCtrl abstract.ClientCtrl) {
	builder.clientCtrl = clientCtrl
}

var (
	Register = func() (string, func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder) {
		Module := func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder {
			newBuilder := new(EchoCommandBuilder)
			newBuilder.SetClientCtrl(clientCtrl)
			return newBuilder
		}
		return "echo", Module
	}
)
