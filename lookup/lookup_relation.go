package lookup

import (
	"fmt"
	"groupie-tracker/apihandlers"
)

type ConcertInfo struct {
	ID		int
	Name	string
	CreationYear	int
	Members	[]string
	DatesLocations	map[string][]string
}

func RelationLookup() ([]ConcertInfo, error) {
	artists, err := handlers.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return nil, err
	}  

	relations, err := handlers.FetchRelations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		return nil, err
	}

	relationsByID := make(map[int]handlers.Relation, len(relations.Index))
	for _, relation := range relations.Index {
		relationsByID[relation.ID] = relation
	}

	concerts := []ConcertInfo{}
	
	for _, artist := range artists {
		relation, ok := relationsByID[artist.ID]
		if ok{
			concert := ConcertInfo {
				ID:				artist.ID,
				Name:			artist.Name,
				CreationYear:	artist.CreationYear,
				Members:		artist.Members,
				DatesLocations:	relation.DatesLocations,
			}
			
			concerts = append(concerts, concert)

		}else{
			fmt.Println("No concert information found")
		}
	}
	return concerts, nil
}
