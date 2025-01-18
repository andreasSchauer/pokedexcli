package pokeapi

import (
	"encoding/json"
	"fmt"
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
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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
	if res.StatusCode > 299 {
		return LocationsList{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	locationsList := LocationsList{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		return LocationsList{}, err
	}

	return locationsList, nil
}


