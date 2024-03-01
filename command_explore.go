package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location provided")
	}
	location := args[0]

	res, err := cfg.pokeapiServiceClient.GetLocationDetails(location)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	fmt.Printf("Exploring (%d) %v...\n", res.ID, res.Name)
	fmt.Println("Pokemon Found: ")

	for _, l := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", l.Pokemon.Name)
	}
	return nil

}
