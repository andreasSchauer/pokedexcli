package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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


func commandMapf(cfg *config) error {
	locationsList, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.previousLocationsURL = locationsList.Previous
	cfg.nextLocationsURL = locationsList.Next

	locations := locationsList.Results

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	return nil
}


func commandMapb(cfg *config) error {
	locationsList, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.previousLocationsURL = locationsList.Previous
	cfg.nextLocationsURL = locationsList.Next

	locations := locationsList.Results

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	return nil
}
