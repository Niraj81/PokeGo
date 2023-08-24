package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	println("Welcom To PokeGo!")
	println("Usage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s : %s\n", cmd.name, cmd.description)
	}
	println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
