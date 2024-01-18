package internal

import "fmt"

func CreateMessage() string {
	return "Hello, world."
}

func ParseMessage(message string) string {
	return fmt.Sprintf("You said: %s", message)
}
