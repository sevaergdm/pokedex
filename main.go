package main

import (
	"time"

	"github.com/sevaergdm/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	caughtPokemon := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemon,
	}

	startRepl(cfg)
}
