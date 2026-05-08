package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: "",
	}
	for {
		var input string
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		inputWords := cleanInput(input)
		validCommands := getCommands()
		cmd, ok := validCommands[inputWords[0]]
		if ok {
			err := cmd.callback(&conf)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
