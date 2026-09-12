package lookup

import (
	"groupie-tracker/models"
	"strings"
)

func SearchArtists(artists []models.Artist, search string) []models.Artist {
	filteredArtists := []models.Artist{}
	for _, artist := range artists {
		if strings.Contains(
			strings.ToLower(artist.Name),
			strings.ToLower(search)) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}
