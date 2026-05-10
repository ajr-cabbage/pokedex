package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ajr-cabbage/pokedex/internal/pokecache"
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

type LocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	cacheRes, ok := cache.Get(url)
	if ok {
		areas := LocationAreas{}
		err := json.Unmarshal(cacheRes, &areas)
		if err == nil {
			return areas, nil
		} else {
			return areas, err
		}
	}

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

	cache.Add(url, body)

	areas := LocationAreas{}

	err = json.Unmarshal(body, &areas)
	if err != nil {
		return LocationAreas{}, err
	}

	return areas, nil
}

func GetLocationData(url string, cache *pokecache.Cache) (LocationData, error) {
	cacheRes, ok := cache.Get(url)
	if ok {
		locData := LocationData{}
		err := json.Unmarshal(cacheRes, &locData)
		if err == nil {
			return locData, nil
		} else {
			return locData, err
		}
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationData{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationData{}, errors.New("Response failed")
	}
	if err != nil {
		return LocationData{}, err
	}

	cache.Add(url, body)

	locDat := LocationData{}

	err = json.Unmarshal(body, &locDat)
	if err != nil {
		return LocationData{}, err
	}

	return locDat, nil
}
