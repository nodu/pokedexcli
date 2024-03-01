package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	// Below not needed as we'll just go back to the beginning?
	// if cfg.nextLocationAreaURL != nil {
	// 	return errors.New("No more areas, you're on the last page")
	// }

	res, err := cfg.pokeapiServiceClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		// log.Fatal(err)
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
