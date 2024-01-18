package main

import (
	"fmt"

	"github.com/ymtdzzz/go-multimodule-test/internal"
)

func GetMessage() string {
	return fmt.Sprintln(internal.CreateMessage())
}
