package runner

import (
	"encoding/json"
	"fmt"
	"tcp-client-runner/abstract"
	"tcp-client-runner/utils"
	"tcp-client-runner/utils/crypto"
	"tcp-client-runner/utils/io"
	"tcp-client-runner/utils/logger"
)

// Chat Command For send chat message
type ChatCommand struct {
	tcpClient abstract.Client
}

func (command *ChatCommand) Execute(data map[string]string) {
	if !command.tcpClient.IsLogin() {
		logger.Warning("You are not login now!")
		return
	}

	fmt.Println("Say something to other players, Enjoy!")
	content := io.ReadLine("Enter Message", "")

	if len(content) <= 0 {
		return
	}

	command.tcpClient.SendMessage(map[string]interface{}{
		"type": "chat",
		"body": map[string]string{
			"msg": content,
		},
	})
}

func (command *ChatCommand) Fields() []string {
	return []string{}
}

// General Command For general purples
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

// Login Command For user login
type LoginCommand struct {
	tcpClient abstract.Client
}

func (command *LoginCommand) Execute(data map[string]string) {
	if !command.tcpClient.IsConnect() {
		logger.Warning("Please connect server first!")
		return
	}
	if command.tcpClient.IsLogin() {
		logger.Warning("You have Login now!")
		return
	}

	fmt.Println("Please Set Username For Login!")
	username := io.ReadLine("Enter Username", "visitor")
	gameId := io.ReadLine("Enter Game Id", "")
	protocol := io.ReadLine("Enter Protocol", "json")
	uid := utils.GenerateObjectId()

	if len(gameId) == 0 {
		logger.Warning("You need point the game you want")
		return
	}
	command.tcpClient.SetUid(uid)
	command.tcpClient.SetUsername(username)
	command.tcpClient.SetTempProtocol(protocol)

	coder, _ := crypto.GenerateCoder("json")
	responseStr, _ := coder.Encode(map[string]interface{}{
		"type": "login",
		"body": map[string]interface{}{
			"gameId":   gameId,
			"username": username,
			"protocol": protocol,
			"uid":      command.tcpClient.GetUid(),
			"token":    "",
		},
	})
	command.tcpClient.SendBytes(responseStr)
	logger.Warning(username + ":" + protocol)

	command.tcpClient.Login(true)
}
func (command *LoginCommand) Fields() []string {
	return []string{}
}

// Connect Command For Connect with Tcp Server
type ConnectCommand struct {
	commander *GameCommander
}

func (command *ConnectCommand) Execute(data map[string]string) {
	fmt.Println("Before everything, You have to set remote hostname and server port")

	hostname := io.ReadLine("Enter hostname", "127.0.0.1")
	port := io.ReadLine("Enter port", "50000")
	tcpClient := buildTcpClient(hostname, port)

	command.commander.installClient(tcpClient)
	tcpClient.Connect()
}
func (command *ConnectCommand) Fields() []string {
	return []string{}
}

type HelpCommand struct{}

func (command *HelpCommand) Execute(data map[string]string) {
	fmt.Println("This is a simple TCP client as long live chat app")
	fmt.Println("============  All Command Support Not  ============")
	fmt.Println()
	fmt.Println("Connet with server is necessary")
	fmt.Println()

	fmt.Println("Commands:")
	fmt.Println("  login:    login a game by id")
	fmt.Println("  chat:     send message to peers")
	fmt.Println("  .exit:    quit")
	fmt.Println()
	fmt.Println("More application will be added.")
}

func (command *HelpCommand) Fields() []string {
	return []string{}
}
