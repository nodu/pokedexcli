package main

import "github.com/nodu/pokedexcli/internal/pokeapiService"

type config struct {
	pokeapiServiceClient pokeapiService.Client
	nextLocationAreaURL  *string //url.URL?
	prevLocationAreaURL  *string //url.URL?
}

func main() {
	cfg := config{
		// init client here so it's reused
		// pokeapiServiceClient: pokeapiService.NewClient(5 * time.Second),
		pokeapiServiceClient: pokeapiService.NewClient(),
	}

	startRepl(&cfg)
}
