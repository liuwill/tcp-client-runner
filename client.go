package runner

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"tcp-client-runner/utils/codec"
	"tcp-client-runner/utils/crypto"
	"tcp-client-runner/utils/logger"
	"time"
)

const writeWait = 10 * time.Second

type TcpClient struct {
	hostname      string
	port          string
	username      string
	uid           string
	connection    net.Conn
	protocol      string
	tempProtocol  string
	connectStatus bool
	loginStatus   bool

	message       chan []byte
	quitSemaphore chan bool
}

func (tcpClient *TcpClient) IsConnect() bool {
	return tcpClient.connectStatus
}
func (tcpClient *TcpClient) IsLogin() bool {
	return tcpClient.loginStatus
}
func (tcpClient *TcpClient) Login(loginStatus bool) {
	tcpClient.loginStatus = loginStatus
}
func (tcpClient *TcpClient) SetUid(uid string) {
	tcpClient.uid = uid
}
func (tcpClient *TcpClient) SetUsername(username string) {
	tcpClient.username = username
}
func (tcpClient *TcpClient) SetProtocol(protocol string) {
	tcpClient.protocol = protocol
}
func (tcpClient *TcpClient) SetTempProtocol(protocol string) {
	tcpClient.tempProtocol = protocol
}
func (tcpClient *TcpClient) Connect() {
	if tcpClient.connectStatus {
		return
	}

	var tcpAddr *net.TCPAddr
	var remoteAddr = fmt.Sprintf("%s:%s", tcpClient.hostname, tcpClient.port)
	tcpAddr, resolveErr := net.ResolveTCPAddr("tcp", remoteAddr)

	if resolveErr != nil {
		logger.Error("can't find remote peer!")
		return
	}
	conn, dialErr := net.DialTCP("tcp", nil, tcpAddr)
	tcpClient.connection = conn
	if dialErr != nil {
		logger.Error("connect fail!")
		return
	}
	logger.Success("connected!")

	go onCloseConnect(tcpClient)
	go onMessageReceived(tcpClient)

	tcpClient.connectStatus = true
}

func (tcpClient *TcpClient) Close() {
	tcpClient.quitSemaphore <- true
}

func onCloseConnect(tcpClient *TcpClient) {
	<-tcpClient.quitSemaphore

	logger.Error("disconnected!")
	if tcpClient.connectStatus {
		tcpClient.connection.Close()
		tcpClient.connectStatus = false
		tcpClient.loginStatus = false
	}
}

func onMessageReceived(tcpClient *TcpClient) {
	defer func() {
		tcpClient.quitSemaphore <- true
	}()

	reader := bufio.NewReader(tcpClient.connection)
	buffer := make([]byte, 2048)
	for {
		targetBytes, err := codec.Unpack(reader)

		for i, v := range targetBytes {
			buffer[i] = v
		}

		length := len(targetBytes)
		// length, err := reader.Read(buffer) //.ReaString('\n')
		msg := buffer[0:length]
		// fmt.Println(msg)
		if err != nil {
			tcpClient.quitSemaphore <- true
			break
		}

		if tcpClient.protocol == "protobuf" {
			var rawData interface{}

			coder, _ := crypto.GenerateCoder(tcpClient.protocol)
			_ = coder.Decode(msg, &rawData)
			targetMsg, _ := json.Marshal(rawData)
			tcpClient.message <- targetMsg
		} else {
			var rawData interface{}
			coder, _ := crypto.GenerateCoder("json")
			_ = coder.Decode(msg, &rawData)

			baseData, _ := rawData.(map[string]interface{})
			rawType, _ := baseData["type"]
			baseType, _ := rawType.(string)

			if baseType == "join" {
				tcpClient.SetProtocol(tcpClient.tempProtocol)
			}

			tcpClient.message <- msg
		}
	}
}

func (tcpClient *TcpClient) SendBytes(message []byte) {
	tcpClient.Connect()

	tcpClient.connection.SetWriteDeadline(time.Now().Add(writeWait))

	targetBytes, err := codec.Pack(message)
	if err != nil {
		return
	}
	tcpClient.connection.Write(targetBytes)
}

func (tcpClient *TcpClient) SendMessage(message interface{}) {
	tcpClient.Connect()

	coder, _ := crypto.GenerateCoder(tcpClient.protocol)
	responseStr, _ := coder.Encode(message)

	tcpClient.connection.SetWriteDeadline(time.Now().Add(writeWait))
	targetBytes, err := codec.Pack(responseStr)
	if err != nil {
		return
	}
	tcpClient.connection.Write(targetBytes)
}

func buildTcpClient(hostname string, port string) *TcpClient {
	return &TcpClient{
		hostname:      hostname,
		port:          port,
		connection:    nil,
		connectStatus: false,
		loginStatus:   false,
		protocol:      "json",
		message:       make(chan []byte),
		quitSemaphore: make(chan bool),
	}
}
