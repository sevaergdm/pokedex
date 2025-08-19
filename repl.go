package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sevaergdm/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned_input := cleanInput(input)
		if len(cleaned_input) == 0 {
			continue
		}

		command := cleaned_input[0]
		commandsRegistry := getCommands()
		command_struct, ok := commandsRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command_struct.callback(cfg, cleaned_input[1:]...)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}

func cleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words := strings.Fields(lowered_text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the pokemon at a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attept to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
