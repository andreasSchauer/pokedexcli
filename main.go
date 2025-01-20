package main

import (
	"time"

	"github.com/andreasSchauer/pokedexcli/internal/pokeapi"
	
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	initNextLocationsURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	cfg := &config{
		pokeapiClient: pokeClient,
		nextLocationsURL: &initNextLocationsURL,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}