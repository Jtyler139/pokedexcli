package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	newConfig := Config{Next: "", Previous: nil}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		clean := cleanInput(reader.Text())
		if len(clean) == 0 {
			continue
		}

		commandName := clean[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&newConfig)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
		
	}
}

func cleanInput(text string) []string {
	var split []string
	words := strings.Fields(text)
	for i := range words {
		split = append(split, strings.ToLower(words[i]))
	}
	return split
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		 "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:	 commandMapb,
		},
	}
}

type Page struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Next 		string
	Previous 	*string
}