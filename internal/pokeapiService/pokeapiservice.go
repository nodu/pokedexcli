package pokeapiService

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct { // Create a new client so we can modify the http client globally to set a timeout
	httpClient http.Client
}

// func NewClient(timeout time.Duration) Client { // could pass in a timeout too
func NewClient() Client { // kinda like a constructor
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
			//Timeout: timeout
		},
	}
}
