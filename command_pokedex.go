package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokeName := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokeName)
	}
	return nil
}