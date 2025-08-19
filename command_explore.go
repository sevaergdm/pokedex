package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide a location area")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	locationResponse, err := cfg.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}
	
	fmt.Println("Found Pokemon: ")	
	for _, pokemon := range locationResponse.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}
	return nil
} 
