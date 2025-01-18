package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAreaBatch(url string) (AreaBatch, error) {
	res, err := http.Get(url)
	if err != nil {
		return AreaBatch{}, err
	}
	
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaBatch{}, err
	}
	if res.StatusCode > 299 {
		return AreaBatch{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	areaBatch := AreaBatch{}
	err = json.Unmarshal(body, &areaBatch)
	if err != nil {
		return AreaBatch{}, err
	}

	return areaBatch, nil
}


type AreaBatch struct {
	Count    int    	`json:"count"`
	Next     string 	`json:"next"`
	Previous string    	`json:"previous"`
	Results  []struct {
		Name string 	`json:"name"`
		URL  string 	`json:"url"`
	} 					`json:"results"`
}