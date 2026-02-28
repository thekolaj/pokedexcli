package main

import (
	"time"

	"github.com/thekolaj/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}

type config struct {
	pokeapiClient    pokeapi.Client
	caughtPokemon    map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}
