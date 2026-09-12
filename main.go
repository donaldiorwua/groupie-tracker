package main

import (
	"fmt"
	"groupie-tracker/webhandlers"
	"net/http"
)

func main() {

	fmt.Println("server running at http://localhost:8080/")

	http.HandleFunc("/", webhandlers.IndexHandler)
	http.HandleFunc("/artists", webhandlers.Artists)
	http.HandleFunc("/artist", webhandlers.ArtistProfile)

	http.ListenAndServe(":8080", nil)
}
