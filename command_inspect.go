package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	fmt.Println("Opening the Pokedex...")
	if len(args) == 0 {
		fmt.Printf("You've caught %d Pokemon!\n", len(cfg.pokedex.data))
		for k := range cfg.pokedex.data {
			fmt.Printf("- %s\n", k)
		}

		return nil
	}

	pokemon, caught := cfg.pokedex.data[args[0]]
	if !caught {
		return errors.New("Not Found! Trainer: Get back out there!")
	}

	fmt.Printf("Details on %s:\n", pokemon.Name)
	fmt.Printf("- Height: %v\n", pokemon.Height)
	fmt.Printf("- Weight: %v\n", pokemon.Weight)
	fmt.Printf("- Stats: %v\n", pokemon.Stats)
	fmt.Printf("- Types %v\n", pokemon.Types)

	return nil
}
