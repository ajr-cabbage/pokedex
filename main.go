package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
			err := cmd.callback()
			if err != nil {
				fmt.Errorf("Error: %w", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
