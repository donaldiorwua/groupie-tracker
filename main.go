package main

import (
	"fmt"
	"groupie-tracker/webhandlers"
	"net/http"
)

func main() {
	//var err error

	//webhandlers.Temp = template.Must(template.ParseGlob("templates/*html"))

	fmt.Println("server running at http://localhost:8080/")

	//http.HandleFunc("/", ArtistProfile)
	http.HandleFunc("/artists", webhandlers.Artists)

	http.ListenAndServe(":8080", nil)
}
