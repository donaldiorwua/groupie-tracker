package lookup

import (
	"fmt"
	"errors"
	"groupie-tracker/apihandlers"
	"groupie-tracker/models"
)

var ErrArtistNotFound = errors.New("Artist not found")

func FindArtistProfile(id int) (models.ConcertInfo, error) {
	artists, err := apihandlers.FetchArtists()
	if err != nil {
		return models.ConcertInfo{}, err
	}

	found := false
	var selectedArtist models.Artist
	for _, artist := range artists {
		if artist.ID == id {
			selectedArtist = artist
			found = true
			break
		}
	}

	if !found {
		return models.ConcertInfo{}, ErrArtistNotFound
	}
	

	relations, err := apihandlers.FetchRelations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return models.ConcertInfo{}, err
	}

	relationsByID := make(map[int]models.Relation, len(relations.Index))
	for _, relation := range relations.Index {
		relationsByID[relation.ID] = relation
	}

	var artistProfile models.ConcertInfo

	relation, ok := relationsByID[selectedArtist.ID]
	if ok {
		artistProfile = models.ConcertInfo{
			ID:             selectedArtist.ID,
			Image:          selectedArtist.Image,
			Name:           selectedArtist.Name,
			CreationYear:   selectedArtist.CreationYear,
			FirstAlbum:     selectedArtist.FirstAlbum,
			Members:        selectedArtist.Members,
			DatesLocations: relation.DatesLocations,
		}
	} else {
		return models.ConcertInfo{}, fmt.Errorf("No Artist information found")
		
	}
	
	return artistProfile, nil
}
