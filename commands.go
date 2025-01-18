package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name			string
	description		string
	callback		func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Explanation of all available commands",
			callback:		commandHelp,
		},
	}
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}