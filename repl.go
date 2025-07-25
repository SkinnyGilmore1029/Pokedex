package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")     //print command prompt
		input.Scan()                // get the user iput
		text := input.Text()        // convert and save input to text
		cleaned := cleanInput(text) // cleans the text

		//check to make sure the command is not empty
		if len(cleaned) == 0 {
			continue
		}
		// get the first word of the input
		first_word := cleaned[0]

		// check if the command exists
		if commands, exists := Commands[first_word]; exists {
			if err := commands.callback(); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	var cleaned []string
	noWhitespace := strings.TrimSpace(text)
	lowered := strings.ToLower(noWhitespace)
	splitWords := strings.Fields(lowered)
	cleaned = splitWords

	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
