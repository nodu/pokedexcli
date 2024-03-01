package main

import (
	"fmt"
)

func commandHelp(cfg *config, location string) error { // NOTE help and exit shouldn't have config? // Seems instructor uses this...
	fmt.Println()
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your availble commands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("  - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
