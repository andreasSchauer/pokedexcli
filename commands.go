package main

import (
	"fmt"
	"os"
	"errors"
)


type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:		 "explore",
			description: "List all pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:		 "catch",
			description: "Attempt to catch a pokemon",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect",
			description: "Get data of a caught pokemon",
			callback:	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Display a list of caught pokemon",
			callback:	 commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}


func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}


func commandMapf(cfg *config, args ...string) error {
	if cfg.nextLocationsURL == nil {
		fmt.Println("You're already on the last page")
		fmt.Println()
	}

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


func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		fmt.Println("You're already on the first page")
		fmt.Println()
	}

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


func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]

	pokemonList, err := cfg.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}

	pokemonEncounters := pokemonList.PokemonEncounters
	fmt.Println("Found Pokemon:")

	for _, pkmn := range pokemonEncounters {
		fmt.Printf(" - %s\n", pkmn.Pokemon.Name)
	}

	return nil
}


func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if pokemon.CatchPokemon() {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	
	return nil
}


func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, registered := cfg.pokedex[pokemonName]
	if !registered {
		return fmt.Errorf("you need to catch %s first, before you can inspect it", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Types:\n")

	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	fmt.Printf("Base Stats:\n")

	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", s.Stat.Name ,s.BaseStat)
	}

	return nil
}


func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokedex

	if len(pokedex) == 0 {
		return errors.New("your pokedex is empty")
	}

	fmt.Println("Your Pokedex:")

	for key := range pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}