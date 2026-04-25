package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

type PokeClientConfig struct {
	ClientTimeout time.Duration
	CacheInterval time.Duration
}

// NewClient -
func NewClient(cfg PokeClientConfig) *Client {
	client := Client{
		httpClient: http.Client{
			Timeout: cfg.ClientTimeout,
		},
		cache: pokecache.NewCache(cfg.CacheInterval),
	}
	return &client
}
