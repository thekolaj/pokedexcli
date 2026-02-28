package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemon, ok := cfg.caughtPokemon[args[1]]

	if !ok {
		return fmt.Errorf("you haven't caught %s yet", args[1])
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
