package main

func main() {

	startRepl()

}

/*
Notes from the lessons
-----------------------

You'll want to use functions from Go's strings package. Some helpful ones to look into:

strings.TrimSpace() - removes leading/trailing whitespace
strings.Fields() - splits on whitespace and handles multiple spaces nicely
strings.ToLower() - converts to lowercase

if first_word == "exit" {
			if err := Commands[first_word].callback(); err != nil {
				fmt.Println("Error:",err)
			}
		}

var Commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description:  "Displays a help message",
		callback:    commandHelp,
	},
}
*/
