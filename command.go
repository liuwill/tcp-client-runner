package runner

import (
	"encoding/json"
	"fmt"
	"tcp-client-runner/utils"
	"tcp-client-runner/utils/crypto"
	"tcp-client-runner/utils/logger"
)

type ChatCommand struct {
	tcpClient *TcpClient
}

func (command *ChatCommand) Execute(data map[string]string) {
	if !command.tcpClient.IsLogin() {
		logger.Warning("You are not login now!")
		return
	}

	fmt.Println("Say something to other players, Enjoy!")
	content := ReadLine("Enter Message", "")

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

type LoginCommand struct {
	tcpClient *TcpClient
}

func (command *LoginCommand) Execute(data map[string]string) {
	if command.tcpClient.IsLogin() {
		logger.Warning("You have Login now!")
		return
	}

	fmt.Println("Please Set Username For Login!")
	username := ReadLine("Enter Username", "visitor")
	gameId := ReadLine("Enter Game Id", "")
	protocol := ReadLine("Enter Protocol", "json")
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
			"uid":      command.tcpClient.uid,
			"token":    "",
		},
	})
	command.tcpClient.SendBytes(responseStr)
	logger.Warning(username + ":" + protocol)

	command.tcpClient.Login(true)
	// bytes, _ := json.Marshal(data)
	// logger.Info(string(bytes))
}
func (command *LoginCommand) Fields() []string {
	return []string{}
}

type ConnectCommand struct {
	commander *GameCommander
}

func (command *ConnectCommand) Execute(data map[string]string) {
	fmt.Println("Before everything, You have to set remote hostname and server port")

	hostname := ReadLine("Enter hostname", "127.0.0.1")
	port := ReadLine("Enter port", "50000")
	tcpClient := buildTcpClient(hostname, port)

	command.commander.installClient(tcpClient)
	tcpClient.Connect()
}
func (command *ConnectCommand) Fields() []string {
	return []string{}
}
