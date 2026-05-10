package main

import (
	"fmt"
	"os"
	"strings"

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
