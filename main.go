package main

import (
	"groupie-tracker/webhandlers"
	"log"
	"net/http"
	"os"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", webhandlers.IndexHandler)
	http.HandleFunc("/artists", webhandlers.Artists)
	http.HandleFunc("/artist", webhandlers.ArtistProfile)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	http.ListenAndServe("0.0.0.0:"+port, nil)
}
