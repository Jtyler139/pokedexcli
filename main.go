package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 
		line := scanner.Text()
		clean := cleanInput(line)
		fmt.Printf("Your command was: %s\n", clean[0])
	}
}