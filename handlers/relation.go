package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func FetchRelations() (Relation, error) {
	relations := Relation{}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return Relation{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Relation{}, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}

	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&relations); err != nil {
		return Relation{}, err
	}
	return relations, nil
}

