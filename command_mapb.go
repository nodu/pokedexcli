package main

import (
	"errors"
	"fmt"
)

func callbackMapb(cfg *config, location string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("Can't go back, you're on the first page")
	}
	res, err := cfg.pokeapiServiceClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		// log.Fatal(err) // This far down the callstack don't want to return fatals,
		// best for close to the beginning of start of app, ie init
		return err
	}
	fmt.Println("Location areas: ")

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	for _, l := range res.Results {
		fmt.Printf(" - %s\n", l.Name)
	}
	return nil

}
