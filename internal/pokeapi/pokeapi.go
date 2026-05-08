package "pokeapi"

import (
	"json"
	"http"
	"io"
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

func GetLocationAreas(url) LocationAreas, error{
	res, err := Get(url)
	if err != nil {
		return struct{}, err
	}
	
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return struct{}, errors.New("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return struct{}, err
	}
	
	areas := LocationAreas{}
	
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return struct{}, err
	}
	
	return areas, nil	
}


