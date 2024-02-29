package main

import (
	"fmt"
)

func commandHelp() error {
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
