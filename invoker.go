package runner

import (
	"strings"
)

type Invoker struct {
	command Command
}

func (invoker *Invoker) SetCommand(command Command) {
	invoker.command = command
}

func (invoker *Invoker) Action(input string) {
	invokerData := invoker.Parse(input, invoker.command)
	invoker.command.Execute(invokerData)
}
func (invoker *Invoker) Parse(input string, command Command) map[string]string {
	input = strings.Replace(input, "  ", " ", -1)
	inputPiece := strings.Split(input, " ")
	result := make(map[string]string)

	for index, field := range command.Fields() {
		if index < len(inputPiece) && len(inputPiece[index]) > 0 {
			result[field] = inputPiece[index]
		}
	}

	if len(command.Fields()) == 0 {
		result["content"] = strings.Join(inputPiece[1:], " ")
		return result
	}

	return result
}
