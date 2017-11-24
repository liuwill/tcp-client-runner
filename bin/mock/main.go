package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
}

func (tcpConnection *TcpConnection) Close() {

}

func (tcpConnection *TcpConnection) Read(buffer []byte) (int, error) {
	tcpConnection.conn.SetReadDeadline(time.Now().Add(writeWait))
	targetBytes, err := codec.Unpack(tcpConnection.reader)

	for i, v := range targetBytes {
		buffer[i] = v
	}
	return len(targetBytes), err
}

func (tcpConnection *TcpConnection) Run() {
	defer func() {
		tcpConnection.Close()
	}()
}

type MockServer struct {
	Peers []*TcpConnection
}

func (mockServer *MockServer) handleTcpConn(connection net.Conn) {
	tcpConnection := &TcpConnection{
		conn:   connection,
		reader: bufio.NewReader(connection),
	}
	mockServer.Peers = append(mockServer.Peers, tcpConnection)

	logger.Success(fmt.Sprintf("Connect From: %s", connection.RemoteAddr().String()))
	go startTransform(tcpConnection)
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

	mockServer := MockServer{
		Peers: []*TcpConnection{},
	}
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
