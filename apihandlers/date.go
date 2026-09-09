package apihandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"groupie-tracker/models"
)

func FetchDates() (models.Date, error) {
	dates := models.Date{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return models.Date{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Date{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&dates); err != nil {
		return models.Date{}, err
	}
	return dates, nil
}
