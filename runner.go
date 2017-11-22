package runner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadWord(tip string, def string) string {
	if len(def) > 0 {
		tip = fmt.Sprintf("%s (%s): ", tip, def)
	}

	var input string
	fmt.Print(tip)
	fmt.Scanln(&input)
	if len(input) == 0 {
		return def
	}
	return strings.TrimSpace(input)
}

func ReadLine(tip string, def string) string {
	if len(def) > 0 {
		tip = fmt.Sprintf("%s (%s): ", tip, def)
	} else {
		tip = fmt.Sprintf("%s : ", tip)
	}

	var input string
	fmt.Print(tip)
	// fmt.Scanln(&input)

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return def
	}

	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return def
	}
	return input
}

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
