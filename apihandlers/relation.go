package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"groupie-tracker/models"

)

func FetchRelations() (models.Relations, error) {
	relations := models.Relations{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return models.Relations{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Relations{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&relations); err != nil {
		return models.Relations{}, err
	}
	return relations, nil
}

