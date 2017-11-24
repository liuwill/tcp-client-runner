package main

import (
	runner "tcp-client-runner"
	createBuilder "tcp-client-runner/builder/create"
	echoBuilder "tcp-client-runner/builder/echo"
)

func main() {
	commandRunner := runner.CreateRunner()
	commandRunner.InstallCommands(createBuilder.Register, echoBuilder.Register)

	commandRunner.Bootstrap()
}
