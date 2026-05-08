package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreas{}, errors.New("Response failed")
	}
	if err != nil {
		return LocationAreas{}, err
	}

	areas := LocationAreas{}

	err = json.Unmarshal(body, &areas)
	if err != nil {
		return LocationAreas{}, err
	}

	return areas, nil
}
