package main


import(
	"fmt"
	"groupie-tracker/handlers"
)

func main() {
	artists, err := handlers.FetchArtists()
	if err != nil {
		fmt.Println("Error fetching artists:", err)
		return
	}

	for _, artist := range artists {
		fmt.Printf("ID: %d, Name: %s, Creation Year: %d, Members: %v\n", artist.ID, artist.Name, artist.CreationYear, artist.Members)
	}
}