package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	CreationYear int      `json:"creation_year"`
	Members      []string `json:"members"`
}

func FetchArtists() ([]Artist, error) {
	artists := []Artist{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}
