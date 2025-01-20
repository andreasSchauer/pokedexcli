package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"math/rand"
	"fmt"
)


type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Height  int    `json:"height"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}


func (p Pokemon) CatchPokemon() (success bool) {
	fmt.Printf("Throwing a Pokeball at %s...\n", p.Name)

	catchRate := int(100 - (100 * float64(p.BaseExperience) / 750))

	catchPercentage := rand.Intn(100)

	return catchPercentage <= catchRate
}


func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	cachedData, inCache := c.cache.Get(url)
	if inCache {
		pokemon := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	
	}
	
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}