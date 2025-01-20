package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreasSchauer/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeapiClient        pokeapi.Client
	previousLocationsURL *string
	nextLocationsURL     *string
}


func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}
		
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Printf("Command %s is not a valid command.\n", commandName)
		continue
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	words := strings.Fields(textLower)
	return words
}
