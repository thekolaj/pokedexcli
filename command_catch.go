package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) < 40 {
		fmt.Println(pokemon.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}
	return nil
}
