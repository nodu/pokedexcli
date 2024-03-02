package main

import (
	"github.com/nodu/pokedexcli/internal/pokeapiService"
	"time"
)

type config struct {
	pokeapiServiceClient pokeapiService.Client
	nextLocationAreaURL  *string
	prevLocationAreaURL  *string
	pokedex              map[string]pokeapiService.PokemonResonse
}

func main() {
	cfg := config{
		// init NewClient here so it's reused
		pokeapiServiceClient: pokeapiService.NewClient(5*time.Second, 5*time.Minute),
		pokedex:              make(map[string]pokeapiService.PokemonResonse),
	}
	startRepl(&cfg)
}
