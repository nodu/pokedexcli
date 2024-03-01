package pokeapiService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationDetails(pageURL *string, location string) (LocationAreasDetailResponse, error) {
	fullURL := baseURL + "/location-area/" + location

	cachedData, found := c.cache.Get(fullURL)
	if found {
		loc := LocationAreasDetailResponse{}
		err := json.Unmarshal(cachedData, &loc)
		if err != nil {
			return LocationAreasDetailResponse{}, err
		}

		return loc, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasDetailResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasDetailResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasDetailResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasDetailResponse{}, err
	}

	loc := LocationAreasDetailResponse{}
	err = json.Unmarshal(body, &loc)

	if err != nil {
		return LocationAreasDetailResponse{}, err
	}
	c.cache.Add(fullURL, body)
	return loc, nil
}
