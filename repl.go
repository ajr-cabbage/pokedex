package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ajr-cabbage/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
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

	return commands
}

// returns lowered and stripped array of commands
func cleanInput(text string) []string {
	cleanedSlice := strings.Fields(strings.ToLower(text))
	return cleanedSlice
}

// exit the program
func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display the help message
func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// display next 20 location areas
func commandMap(c *config) error {
	if c.next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(c.next)
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
func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(c.previous)
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
