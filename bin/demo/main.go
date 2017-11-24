package main

import (
	runner "tcp-client-runner"
	createBuilder "tcp-client-runner/builder/create"
)

func main() {
	commandRunner := runner.CreateRunner()
	commandRunner.InstallCommands(createBuilder.Register)

	commandRunner.Bootstrap()
}
