package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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
		location := ""
		if len(cleaned) == 2 {
			location = cleaned[1]
		}

		command, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("...invalid command")
			continue
		}

		err := command.callback(cfg, location) // NOTE code smells because *config and location
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
	callback    func(*config, string) error
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
			name:        "explore",
			description: "Displays pokemon in an area",
			callback:    callbackExplore,
		},
	}
}
