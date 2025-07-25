package main

import (
	"bufio"
	"fmt"
	"os"
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
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		text := input.Text()
		cleaned := cleanInput(text)
		first_word := cleaned[0]
		fmt.Printf("Your command was: %s\n", first_word)
	}

}

/*
Notes from the lessons
-----------------------

You'll want to use functions from Go's strings package. Some helpful ones to look into:

strings.TrimSpace() - removes leading/trailing whitespace
strings.Fields() - splits on whitespace and handles multiple spaces nicely
strings.ToLower() - converts to lowercase
*/
