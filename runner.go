package runner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tcp-client-runner/utils/logger"
)

var inputReader *bufio.Reader
var input string
var err error

type CommandRunner struct {
	builders   map[string]CommandBuilder
	clientCtrl ClientCtrl
}

func CreateRunner() CommandRunner {
	gameCommander := StartGameCommander()
	return CommandRunner{
		clientCtrl: &gameCommander,
		builders:   make(map[string]CommandBuilder),
	}
}

func (runner *CommandRunner) InstallCommands(commandBuilders ...func() (string, func(clientCtrl ClientCtrl) CommandBuilder)) {
	for _, v := range commandBuilders {
		entrance, commandBuilder := v()
		runner.builders[entrance] = commandBuilder(runner.clientCtrl)
	}
}

func (runner *CommandRunner) Bootstrap() {
	fmt.Println("Welcome TCP Client (Enter '.exit' to Quit)🍺")

	inputReader := bufio.NewReader(os.Stdin)
	defer func() {
		logger.Info("Exit! Good bye")
	}()

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
			fmt.Println("Enter 'help' for more info!")
			continue
		}
		// fmt.Printf(input)

		var command Command
		commandBuilder, ok := runner.builders[cmdStr]
		if ok {
			command = commandBuilder.Build()
		} else {
			command = runner.clientCtrl.CreateCommand(cmdStr)
			if command == nil {
				logger.Error("command not found")
				continue
			}
		}

		invoker := Invoker{}
		invoker.SetCommand(command)
		invoker.Action(cmdStr)
	}
}
