package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: explore <location>")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
