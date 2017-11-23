package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	runner "tcp-client-runner"
	"tcp-client-runner/utils/logger"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	fmt.Println("Welcome TCP Client (Enter '.exit' to Quit)🍺")

	inputReader := bufio.NewReader(os.Stdin)
	defer func() {
		logger.Info("Exit! Good bye")
	}()

	gameCommander := runner.StartGameCommander()
	for {
		fmt.Print("> ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			logger.Error("read error")
		}
		cmdStr := strings.TrimSpace(input)

		if cmdStr == ".exit" {
			return
		} else if len(cmdStr) == 0 {
			continue
		}
		// fmt.Printf(input)

		command := gameCommander.CreateCommand(cmdStr)
		if command == nil {
			logger.Error("command not found")
			continue
		}

		invoker := runner.Invoker{}
		invoker.SetCommand(command)
		invoker.Action(cmdStr)
	}
}
