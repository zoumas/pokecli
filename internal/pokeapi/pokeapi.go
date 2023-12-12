package pokeapi

import (
	"net/http"
	"time"

	"github.com/zoumas/pokecli/internal/cache"
)

const LocationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/"

type Client struct {
	httpClient     *http.Client
	cache          *cache.Cache
	Next, Previous *string
}

func NewClient(c *http.Client, cacheReapInterval time.Duration) *Client {
	return &Client{
		httpClient: c,
		cache:      cache.NewCache(cacheReapInterval),
	}
}
