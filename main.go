package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var cleaned []string
	noWhitespace := strings.TrimSpace(text)
	lowered := strings.ToLower(noWhitespace)
	splitWords := strings.Fields(lowered)
	cleaned = splitWords

	return cleaned
}

func main() {
	fmt.Println("Hello, World!")

}

/*
You'll want to use functions from Go's strings package. Some helpful ones to look into:

strings.TrimSpace() - removes leading/trailing whitespace
strings.Fields() - splits on whitespace and handles multiple spaces nicely
strings.ToLower() - converts to lowercase
*/
