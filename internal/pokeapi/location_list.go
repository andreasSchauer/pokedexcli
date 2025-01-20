package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


type LocationsList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func (c *Client) ListLocations(pageURL *string) (LocationsList, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, inCache := c.cache.Get(url)
	if inCache {
		locationsList := LocationsList{}
		err := json.Unmarshal(cachedData, &locationsList)
		if err != nil {
			return LocationsList{}, err
		}

		return locationsList, nil
	
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationsList{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsList{}, err
	}

	c.cache.Add(url, body)

	locationsList := LocationsList{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		return LocationsList{}, err
	}

	return locationsList, nil
}


