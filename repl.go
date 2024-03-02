package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	fmt.Println("c pikachu")
	fmt.Println("c bulbasaur")
	fmt.Println("i pikachu")
	fmt.Println("i bulbasaur")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		cleaned := cleanInput(scanner.Text())

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0] //safe index operation, bc len check
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("...invalid command")
			continue
		}

		err := command.callback(cfg, args...) // NOTE code smells because *config and location
		// but not needed on all commands
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display locations, 20 at a time. Call again for next 20",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays locations, Previous 20",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location}",
			description: "Displays pokemon in an area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Throw a pokeball at a pokemon!",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex {pokemon*optional}",
			description: "Open your Pokedex! If pokemon, then show stats.",
			callback:    callbackPokedex,
		},
		"c": {
			name:        "c {pokemon}",
			description: "Throw a pokeball at a pokemon!",
			callback:    callbackCatch,
		},
		"p": {
			name:        "p {pokemon*optional}",
			description: "Open your Pokedex! If pokemon, then show stats.",
			callback:    callbackPokedex,
		},
	}
}
