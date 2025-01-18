package main

import (
	
)

func main() {
	cfg := &config{
			previous: "",
			next: "https://pokeapi.co/api/v2/location-area/",
		}

	startRepl(cfg)
}