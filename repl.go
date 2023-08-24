package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	println("PokeGo ")

	for {
		print("> ")
		scanner.Scan()

		input := cleanUpInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Print(err)
			}
			continue
		} else {
			fmt.Println("Unknow Command")
		}

	}
}

func cleanUpInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
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
			description: "Exit the PokeGo",
			callback:    commandExit,
		},
	}
}
