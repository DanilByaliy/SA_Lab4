package engine

import (
	"strings"
)

func Parse(commandLine string) Command {
	var cmd Command
	parts := strings.Fields(commandLine)

	if strings.TrimSpace(commandLine) == "" {
		return PrintCommand("SYNTAX ERROR: no command")
	}
	if len(parts) < 2 {
		return PrintCommand("SYNTAX ERROR: no argument")
	}

	switch parts[0] {
	case "print":
		cmd = PrintCommand(parts[1])
	case "palindrom":
		cmd = PalindromCommand(parts[1])
	default:
		cmd = PrintCommand("SYNTAX ERROR: invalid command")
	}
	return cmd
}