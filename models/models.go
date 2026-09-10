package models


type ConcertInfo struct {
	ID		int
	Image	string
	Name	string
	CreationYear	int
	FirstAlbum	string
	Members	[]string
	DatesLocations	map[string][]string
}

type Artist struct {
	ID           int      `json:"id"`
	Image		 string	  `json:"image"`
	Name         string   `json:"name"`
	CreationYear int      `json:"creationDate"`
	FirstAlbum	 string	  `json:"firstAlbum"`
	Members      []string `json:"members"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

type Location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Date struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
 