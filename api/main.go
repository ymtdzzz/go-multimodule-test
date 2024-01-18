package main

import (
	"fmt"

	"github.com/ymtdzzz/go-multimodule-test/internal"
)

func main() {
	fmt.Println(internal.CreateMessage())
}

func GetMessage() string {
	return fmt.Sprintln(internal.CreateMessage())
}
