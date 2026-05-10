package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ajr-cabbage/pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := Config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: "",
		cache:    pokecache.NewCache(5 * time.Second),
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
			err := cmd.callback(&conf, inputWords)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
