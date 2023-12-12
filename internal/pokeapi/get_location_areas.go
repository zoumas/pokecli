package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (*LocationAreasResponse, error) {
	data, ok := c.cache.Get(url)
	if !ok {
		resp, err := c.httpClient.Get(url)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
		}

		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		c.cache.Add(url, data)
	}

	v := LocationAreasResponse{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
