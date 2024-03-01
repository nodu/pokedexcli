package main

import (
	"github.com/nodu/pokedexcli/internal/pokeapiService"
	"time"
)

type config struct {
	pokeapiServiceClient pokeapiService.Client
	nextLocationAreaURL  *string
	prevLocationAreaURL  *string
}

func main() {
	cfg := config{
		// init NewClient here so it's reused
		pokeapiServiceClient: pokeapiService.NewClient(5*time.Second, 5*time.Minute),
	}
	startRepl(&cfg)
}
