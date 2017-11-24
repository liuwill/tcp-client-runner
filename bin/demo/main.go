package main

import (
	runner "tcp-client-runner"
)

func main() {
	commandRunner := runner.CreateRunner()
	commandRunner.InstallCommands(runner.CreateCommandBuilderRegister)

	commandRunner.Bootstrap()
}
