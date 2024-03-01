package main

import (
	"fmt"
)

func callbackExplore(cfg *config, location string) error {

	res, err := cfg.pokeapiServiceClient.ListLocationDetails(cfg.nextLocationAreaURL, location)
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
