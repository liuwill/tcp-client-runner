package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"tcp-client-runner/utils"
	"tcp-client-runner/utils/codec"
	"tcp-client-runner/utils/logger"
	"time"
)

const (
	// Time allowed to read from or write a message to the peer.
	writeWait = 30 * time.Second
)

type TcpConnection struct {
	id         string
	conn       net.Conn
	reader     *bufio.Reader
	mockServer *MockServer
	message    chan []byte
}

func (tcpConnection *TcpConnection) Close() {
	tcpConnection.mockServer.outPeerChan <- tcpConnection
}

func (tcpConnection *TcpConnection) Read(buffer []byte) (int, error) {
	// tcpConnection.conn.SetReadDeadline(time.Now().Add(writeWait))
	targetBytes, err := codec.Unpack(tcpConnection.reader)

	for i, v := range targetBytes {
		buffer[i] = v
	}
	return len(targetBytes), err
}

func (tcpConnection *TcpConnection) SendMessage(message []byte) {
	tcpConnection.message <- message
}

func (tcpConnection *TcpConnection) Run() {
	defer func() {
		tcpConnection.Close()
	}()

	for {
		message := <-tcpConnection.message
		tcpConnection.conn.SetWriteDeadline(time.Now().Add(writeWait))
		targetBytes, err := codec.Pack(message)
		if err != nil {
			return
		}
		tcpConnection.conn.Write(targetBytes)
	}
}

type MockServer struct {
	peers       map[string]*TcpConnection
	inPeerChan  chan *TcpConnection
	outPeerChan chan *TcpConnection
}

func NewMockServer() *MockServer {
	mockServer := &MockServer{
		peers:       make(map[string]*TcpConnection),
		inPeerChan:  make(chan *TcpConnection),
		outPeerChan: make(chan *TcpConnection),
	}

	go mockServer.Run()
	return mockServer
}

func (mockServer *MockServer) Run() {
	for {
		select {
		case inPeer := <-mockServer.inPeerChan:
			mockServer.peers[inPeer.id] = inPeer
		case outPeer := <-mockServer.outPeerChan:
			delete(mockServer.peers, outPeer.id)
		}
	}
}

func (mockServer *MockServer) handleTcpConn(connection net.Conn) {
	tcpConnection := &TcpConnection{
		id:         utils.GenerateObjectId(),
		conn:       connection,
		reader:     bufio.NewReader(connection),
		mockServer: mockServer,
		message:    make(chan []byte),
	}

	logger.Success(fmt.Sprintf("Connect From: %s", connection.RemoteAddr().String()))
	go startTransform(tcpConnection)

	mockServer.inPeerChan <- tcpConnection
}

func startTransform(tcpConnection *TcpConnection) {
	defer func() {
		tcpConnection.Close()
	}()

	go tcpConnection.Run()
	buffer := make([]byte, 2048)
	for {
		n, err := tcpConnection.Read(buffer)
		if err != nil {
			log.Println(fmt.Sprintf("Read message error: %s, session will be closed immediately", err.Error()))
			return
		}

		if n <= 0 {
			continue
		}

		logger.Info(string(buffer[0:n]))
		tcpConnection.SendMessage(buffer[0:n])
		// handler.handlerRequest(connSession, buffer, n)
	}
}

func main() {
	defaultPort := os.Getenv("DEFAULT_PORT")
	if len(defaultPort) <= 0 {
		defaultPort = "50000"
	}

	listener, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}

	// 监听并接受来自客户端的连接
	log.Printf(" TCP Mock Server listen at %s 🚀 \n", defaultPort)

	mockServer := NewMockServer()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}

		// 异步处理，防止阻塞处理routine
		mockServer.handleTcpConn(conn)
	}
}
