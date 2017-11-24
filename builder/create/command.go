package runner

import (
	"fmt"
	"tcp-client-runner/abstract"
)

type CreateCommand struct {
	tcpClient abstract.Client
}

func (command *CreateCommand) Execute(data map[string]string) {
	fmt.Println("ok")
}

func (command *CreateCommand) Fields() []string {
	return []string{}
}

type CreateCommandBuilder struct {
	clientCtrl abstract.ClientCtrl
}

func (builder *CreateCommandBuilder) Build() abstract.Command {
	return &CreateCommand{
		tcpClient: builder.clientCtrl.GetClient(),
	}
}

func (builder *CreateCommandBuilder) SetClientCtrl(clientCtrl abstract.ClientCtrl) {
	builder.clientCtrl = clientCtrl
}

var (
	Register = func() (string, func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder) {
		Module := func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder {
			newBuilder := new(CreateCommandBuilder)
			newBuilder.SetClientCtrl(clientCtrl)
			return newBuilder
		}
		return "create", Module
	}
)
