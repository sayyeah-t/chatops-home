package util

import (
	"strings"
)

func ParseCommand(msg string) []string {
	msg = strings.Replace(msg, "　", " ", -1)
	return strings.Split(msg, " ")
}

func IsCommand(command string) bool {
	return strings.HasPrefix(command, "!")
}
