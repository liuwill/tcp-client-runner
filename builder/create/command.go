package runner

import (
	"fmt"
	"tcp-client-runner/abstract"
)

type CreateCommand struct {
	tcpClient *abstract.Client
}

func (command *CreateCommand) Execute(data map[string]string) {
	fmt.Println("ok")
}

func (command *CreateCommand) Fields() []string {
	return []string{}
}

type CreateCommandBuilder struct {
	client *abstract.Client
}

func (builder *CreateCommandBuilder) Build() abstract.Command {
	return &CreateCommand{
		tcpClient: builder.client,
	}
}

func (builder *CreateCommandBuilder) SetClient(client *abstract.Client) {
	builder.client = client
}

var (
	Register = func() (string, func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder) {
		Module := func(clientCtrl abstract.ClientCtrl) abstract.CommandBuilder {
			newBuilder := new(CreateCommandBuilder)
			return newBuilder
		}
		return "create", Module
	}
)
