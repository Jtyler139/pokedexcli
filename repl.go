package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/jtyler139/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeapiClient	 pokeapi.Client
	caughtPokemon	 map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:		 "explore <location name>",
			description: "Explore a location",
			callback:	 commandExplore,
		},
		"map": {
			name:		 "map",
			description: "Displays the next page of locations",
			callback:	 commandMapf,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays the previous page of locations",
			callback:	 commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"catch": {
			name:		 "catch",
			description: "Attempt to catch <pokemon>",
			callback:	 commandCatch,
		},
	}
}