package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	clientConfig := pokeapi.PokeClientConfig{
		ClientTimeout: 5 * time.Second,
		CacheInterval: 5 * time.Second,
	}
	pokeClient := pokeapi.NewClient(clientConfig)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
