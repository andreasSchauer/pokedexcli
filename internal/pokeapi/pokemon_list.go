package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)



type PokemonList struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}



func (c *Client) ListPokemon(locationName string) (PokemonList, error) {
	url := baseURL + "/location-area/" + locationName

	cachedData, inCache := c.cache.Get(url)
	if inCache {
		pokemonList := PokemonList{}
		err := json.Unmarshal(cachedData, &pokemonList)
		if err != nil {
			return PokemonList{}, err
		}

		return pokemonList, nil
	
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonList{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonList{}, err
	}

	c.cache.Add(url, body)

	pokemonList := PokemonList{}
	err = json.Unmarshal(body, &pokemonList)
	if err != nil {
		return PokemonList{}, err
	}

	return pokemonList, nil
}


