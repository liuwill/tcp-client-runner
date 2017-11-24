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
