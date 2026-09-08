package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"groupie-tracker/models"
)


func FetchLocations() (models.Location, error) {
	locations := models.Location{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return models.Location{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Location{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&locations); err != nil {
		return models.Location{}, err
	}
	return locations, nil
}
