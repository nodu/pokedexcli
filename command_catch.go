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

	res, err := cfg.pokeapiServiceClient.GetPokemon(name)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	diceRoll := rand.Intn(res.BaseExperience)
	chance := 50
	fmt.Println("roll", diceRoll, "\n", "chanceToCatch", chance)
	if diceRoll > chance {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	//save pokemon data for later
	cfg.pokedex.data[name] = res
	return nil
}
