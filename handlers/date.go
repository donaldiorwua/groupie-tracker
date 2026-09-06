package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Date struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func FetchDates() (Date, error) {
	dates := Date{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return Date{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Date{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&dates); err != nil {
		return Date{}, err
	}
	return dates, nil
}



