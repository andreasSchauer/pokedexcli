package main

import (
	"fmt"
	"os"
	"github.com/andreasSchauer/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name			string
	description		string
	callback		func(cfg *config) error
}

type config struct {
	previous		string
	next			string
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
		"map": {
			name:			"map",
			description:	"Shows areas in batches of 20 (moves forward)",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Shows areas in batches of 20 (moves backward)",
			callback:		commandMapb,
		},
	}
}


func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(cfg *config) error {
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


func commandMap(cfg *config) error {
	if cfg.next == "" {
		return fmt.Errorf("reached final page")
	}

	areaBatch, err := pokeapi.GetAreaBatch(cfg.next)
	if err != nil {
		return err
	}

	cfg.previous = areaBatch.Previous
	cfg.next = areaBatch.Next
	
	areas := areaBatch.Results

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	
	return nil
}


func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		return fmt.Errorf("reached first page")
	}

	areaBatch, err := pokeapi.GetAreaBatch(cfg.previous)
	if err != nil {
		return err
	}

	cfg.previous = areaBatch.Previous
	cfg.next = areaBatch.Next

	areas := areaBatch.Results

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	
	return nil
}