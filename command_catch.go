package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide a pokemon")
	}

	catchResponse, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", catchResponse.Name)
	baseExperience := catchResponse.BaseExperience
	roll := rand.Intn(601)
	
	if roll >= baseExperience {
		cfg.caughtPokemon[catchResponse.Name] = catchResponse
		fmt.Printf("%s was caught!\n", catchResponse.Name)
	} else {
		fmt.Printf("%s escaped!\n", catchResponse.Name)
	}
	return nil
}
