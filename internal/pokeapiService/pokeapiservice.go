package pokeapiService

import (
	"github.com/nodu/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct { // Create a new client so we can modify the http client globally to set a timeout
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client { // could pass in a timeout too
	// kinda like a constructor
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}
