package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	name := args[0]

	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, sta := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", sta.Stat.Name, sta.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf("- %s\n", typ.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}