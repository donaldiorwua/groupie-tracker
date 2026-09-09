package webhandlers

import (
	"fmt"
	"groupie-tracker/apihandlers"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func ArtistProfile(w http.ResponseWriter, r *http.Request) {
	var Temp *template.Template

	if r.Method == http.MethodGet {
		strID := r.URL.Query().Get("id")

		intID, err := strconv.Atoi(strID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		artists, err := apihandlers.FetchArtists()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		found := false
		var selectedArtist models.Artist
		for _, artist := range artists {
			if artist.ID == intID {
				selectedArtist = artist
				found = true
				break
			}
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Artist not found")
			return
		}
		Temp, err = template.ParseFiles("templates/artist.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		err = Temp.Execute(w, selectedArtist)
		if err != nil {
			log.Println(err)
			return
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}
}
