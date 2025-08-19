package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location

	var responseData []byte
	responseData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}

		c.cache.Add(url, data)

		responseData = data
	}

	var pokemon LocationArea
	err := json.Unmarshal(responseData, &pokemon)
	if err != nil {
		return LocationArea{}, err
	}

	return pokemon, nil
}
