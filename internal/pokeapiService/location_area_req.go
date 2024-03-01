package pokeapiService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	fullURL := baseURL + "/location-area"

	if pageURL != nil {
		fullURL = *pageURL
	}
	// I want cache here?
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	loc := LocationAreasResponse{}
	err = json.Unmarshal(body, &loc)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	return loc, nil
}
