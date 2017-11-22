package runner

import (
	"bufio"
	"fmt"
	"net"
	"tcp-client-runner/utils"
	"tcp-client-runner/utils/crypto"
	"tcp-client-runner/utils/logger"
	"time"
)

const writeWait = 10 * time.Second

type TcpClient struct {
	username      string
	uid           string
	connection    net.Conn
	protocol      string
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
func (tcpClient *TcpClient) Connect() {
	if tcpClient.connectStatus {
		return
	}

	var tcpAddr *net.TCPAddr
	tcpAddr, resolveErr := net.ResolveTCPAddr("tcp", "127.0.0.1:50000")

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
	go onMessageRecived(tcpClient)

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

func onMessageRecived(tcpClient *TcpClient) {
	defer func() {
		tcpClient.quitSemaphore <- true
	}()

	reader := bufio.NewReader(tcpClient.connection)
	for {
		buffer := make([]byte, 2048)
		length, err := reader.Read(buffer) //.ReadString('\n')
		msg := buffer[0:length]
		fmt.Println(msg)
		if err != nil {
			tcpClient.quitSemaphore <- true
			break
		}

		tcpClient.message <- msg
	}
}

func (tcpClient *TcpClient) SendMessage(message interface{}) {
	tcpClient.Connect()

	coder, _ := crypto.GenerateCoder(tcpClient.protocol)
	responseStr, _ := coder.Encode(message)

	tcpClient.connection.SetWriteDeadline(time.Now().Add(writeWait))
	tcpClient.connection.Write(responseStr)
}

func buildTcpClient() *TcpClient {
	return &TcpClient{
		connection:    nil,
		connectStatus: false,
		loginStatus:   false,
		protocol:      "json",
		uid:           utils.GenerateObjectId(),
		message:       make(chan []byte),
		quitSemaphore: make(chan bool),
	}
}
