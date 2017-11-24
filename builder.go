package runner

import (
	"fmt"
)

type CreateCommand struct {
	tcpClient *TcpClient
}

func (command *CreateCommand) Execute(data map[string]string) {
	fmt.Println("ok")
}

func (command *CreateCommand) Fields() []string {
	return []string{}
}

type CreateCommandBuilder struct {
	client *TcpClient
}

func (builder *CreateCommandBuilder) Build() Command {
	return &CreateCommand{
		tcpClient: builder.client,
	}
}

func (builder *CreateCommandBuilder) SetClient(client *TcpClient) {
	builder.client = client
}

var (
	CreateCommandBuilderRegister = func() (string, func(clientCtrl ClientCtrl) CommandBuilder) {
		Module := func(clientCtrl ClientCtrl) CommandBuilder {
			newBuilder := new(CreateCommandBuilder)
			return newBuilder
		}
		return "create", Module
	}
)
