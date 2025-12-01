package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationArea string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", locationArea)
	fmt.Println("Found Pokemon:")
	for _, pok := range pokemonResp.PokemonEncounters {
		fmt.Printf("-%v\n", pok.Pokemon.Name)
	}
	return nil
}