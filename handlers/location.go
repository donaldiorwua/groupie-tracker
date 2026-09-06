package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

func FetchLocations() (Location, error) {
	locations := Location{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return Location{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&locations); err != nil {
		return Location{}, err
	}
	return locations, nil
}
