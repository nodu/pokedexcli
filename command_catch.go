package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided!")
	}
	name := args[0]

	pokemon, err := cfg.pokeapiServiceClient.GetPokemon(name)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	diceRoll := rand.Intn(pokemon.BaseExperience)
	chance := 50
	// fmt.Println("roll", diceRoll, "\n", "chanceToCatch", chance)
	if diceRoll > chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	//save pokemon data for later
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}
