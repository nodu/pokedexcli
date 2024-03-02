package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Opening the Pokedex...")
	if len(args) == 0 {
		fmt.Printf("You've caught %d Pokemon!\n", len(cfg.pokedex))
		for k := range cfg.pokedex {
			fmt.Printf("- %s\n", k)
		}

		return nil
	}

	pokemon, caught := cfg.pokedex[args[0]]
	if !caught {
		return errors.New("Not Found! Trainer: Get back out there!")
	}

	fmt.Printf("Details on %s:\n", pokemon.Name)
	fmt.Printf("- Height: %v\n", pokemon.Height)
	fmt.Printf("- Weight: %v\n", pokemon.Weight)
	fmt.Println("\n Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("\n Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}
