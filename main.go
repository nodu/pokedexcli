package main

import (
	// "fmt"
	"github.com/nodu/pokedexcli/internal/pokeapiService"
	"github.com/nodu/pokedexcli/internal/pokecache"
	"time"
)

type config struct {
	pokeapiServiceClient pokeapiService.Client
	nextLocationAreaURL  *string
	prevLocationAreaURL  *string
	cache                pokecache.Cache
}

func main() {
	cfg := config{
		// init NewClient here so it's reused
		// pokeapiServiceClient: pokeapiService.NewClient(5 * time.Second), // Can add a timeout via passing a duration down
		pokeapiServiceClient: pokeapiService.NewClient(),
		cache:                pokecache.NewCache(2 * time.Second),
	}

	// cache := fmt.Println(cache)
	startRepl(&cfg)
}
