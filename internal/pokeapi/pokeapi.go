package pokeapi

import "net/http"

const LocationAreaEndpoint = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

type Client struct {
	httpClient     *http.Client
	Next, Previous *string
}

func NewClient(c *http.Client) *Client {
	return &Client{
		httpClient: c,
	}
}
