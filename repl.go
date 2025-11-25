package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

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
		fmt.Printf("Your command was: %s\n", clean[0])
	}
}