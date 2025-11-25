package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

var commands = make(map[string]cliCommand)

func cleanInput(text string) []string {
	var split []string
	words := strings.Fields(text)
	for i := range words {
		split = append(split, strings.ToLower(words[i]))
	}
	return split
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		line := scanner.Text()
		clean := cleanInput(line)
		command, exists := commands[clean[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func init() {
	commands["help"] = cliCommand{
		name:			"help",
		description:	"Displays a help message",
		callback:		commandHelp,
	}
	commands["exit"] = cliCommand{
		name:			"exit",
		description:	"Exit the Pokedex",
		callback:		commandExit,
	}
}