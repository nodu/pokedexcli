package pokeapiService

type LocationAreasResponse struct { // Capitalized because we want it to be exported form this package
	Count    int     `json:"count"`
	Next     *string `json:"next"`     //At the end of last paginated URL, there will be no next page
	Previous *string `json:"previous"` //At the beginning, there will be no previous page
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
