package main

import (
	"github.com/nodu/pokedexcli/internal/pokeapiService"
	"time"
)

type config struct {
	pokeapiServiceClient pokeapiService.Client
	nextLocationAreaURL  *string
	prevLocationAreaURL  *string
	pokedex
}

type pokedex struct {
	data map[string]pokeapiService.PokemonResonse
}

func main() {
	pd := pokedex{
		data: make(map[string]pokeapiService.PokemonResonse),
	}
	cfg := config{
		// init NewClient here so it's reused
		pokeapiServiceClient: pokeapiService.NewClient(5*time.Second, 5*time.Minute),
		pokedex:              pd,
	}
	startRepl(&cfg)
}
