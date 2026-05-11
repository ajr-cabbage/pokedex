package main

import (
	"fmt"
	"os"
	"strings"
	"math/rand"

	"github.com/ajr-cabbage/pokedex/internal/pokeapi"
	"github.com/ajr-cabbage/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex	 map[string]pokeapi.PokemonData
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays next 20 location areas",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays previous 20 location areas",
		callback:    commandMapb,
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays list of pokemon in an area",
		callback:    commandExplore,
	}
	
	commands["catch"] = cliCommand{
		name:	     "catch",
		description: "Attempts to catch a pokemon and add it to the pokedex",
		callback:    commandCatch,
	}
	
	commands["inspect"] = cliCommand{
		name:	     "inspect",
		description: "View the stats of pokemon you have caught",
		callback:    commandInspect,
	}
	
	commands["pokedex"] = cliCommand{
		name:	     "pokedex",
		description: "View the contents of your pokedex",
		callback:    commandPokedex,
	}

	return commands
}

// returns lowered and stripped array of commands
func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

// exit the program
func commandExit(c *Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display the help message
func commandHelp(c *Config, params []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// display next 20 location areas
func commandMap(c *Config, params []string) error {
	if c.next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(c.next, c.cache)
	if err != nil {
		return err
	}

	if areas.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *areas.Previous
	}

	if areas.Next == nil {
		c.next = ""
	} else {
		c.next = *areas.Next
	}

	for _, loc := range areas.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

// display prev 20 location areas
func commandMapb(c *Config, params []string) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(c.previous, c.cache)
	if err != nil {
		return err
	}

	if areas.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *areas.Previous
	}

	if areas.Next == nil {
		c.next = ""
	} else {
		c.next = *areas.Next
	}

	for _, loc := range areas.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}
// display list of pokemon in a given area
func commandExplore(c *Config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage: explore <area_name>")
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + params[1]

	locData, err := pokeapi.GetLocationData(url, c.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", params[1])

	for _, encounter := range locData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
// attempt to catch a pokemon
func commandCatch(c *Config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage: catch <pokemon_name>")
		return nil
	}
	
	url := "https://pokeapi.co/api/v2/pokemon/" + params[1]
	
	pokemonData, err := pokeapi.GetPokemonData(url, c.cache)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", params[1])
	
	catchProbability := 1.0 - (float64(pokemonData.BaseExperience)/250.0)
	catchRand := rand.Float64()
	
	if catchRand <= catchProbability {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		c.pokedex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s excaped!\n", pokemonData.Name)
	}
	
	return nil
}
// display stats for a pokemon in your pokedex
func commandInspect(c *Config, params []string) error {
	if len(params) < 2 {
		fmt.Println("Usage: inspect <pokemon_name>")
		return nil
	}
	
	pokemonData, ok := c.pokedex[params[1]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	
	fmt.Printf("Name: %s/n", pokemonData.Name)
	fmt.Printf("Height: %d/n", pokemonData.Height)
	fmt.Printf("Weight: %d/n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonData.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}
	
	return nil
}

func commandPokedex(c *Config, params []string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}
	
	for key, _ := range c.pokedex {
		fmt.Printf(" - %s\n", key)
	}	
	
	return nil
}

