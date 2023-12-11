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

func (c *Client) GetLocationAreas(url string) (LocationAreasResponse, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	if resp.StatusCode >= 400 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	v := LocationAreasResponse{}
	if err := json.Unmarshal(data, &v); err != nil {
		return LocationAreasResponse{}, err
	}

	return v, nil
}
