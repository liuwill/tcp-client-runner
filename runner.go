package runner

import (
	"fmt"
	"strings"
)

func ReadLine(tip string, def string) string {
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
