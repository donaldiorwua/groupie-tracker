package lookup

import (
	"fmt"
	"groupie-tracker/apihandlers"
)

func Location_Lookup() (error) {
	locations, err := apihandlers.FetchLocations()
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return nil
	}

	artists, err := apihandlers.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return nil
	}

	for _, artist := range artists {
		artist_ID := artist.ID

		for _, location := range locations.Index {
			location_ID := location.ID

			if artist_ID == location_ID {
				fmt.Printf("ID: %d Artist Name: %s \nLocation: %v\n", artist.ID, artist.Name, location.Locations)
			}
		}
	}
	return nil
}
