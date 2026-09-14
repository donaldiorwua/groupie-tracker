package lookup

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func SearchArtists(artists []models.Artist, search string) []models.Artist {
	filteredArtists := []models.Artist{}
	searchYear, err := strconv.Atoi(search)
	if err == nil {
		for _, artist := range artists {
			if searchYear == artist.CreationYear {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}
	} else {
		for _, artist := range artists {
			if strings.Contains(
				strings.ToLower(artist.Name),
				strings.ToLower(search)) {
				filteredArtists = append(filteredArtists, artist)
				continue
			}
			for _, member := range artist.Members {
				if strings.Contains(
					strings.ToLower(member),
					strings.ToLower(search)) {
					filteredArtists = append(filteredArtists, artist)
					break
				}
			}
		}
	}
	return filteredArtists
}
