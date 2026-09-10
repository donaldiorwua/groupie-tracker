package webhandlers

import (
	"errors"
	"fmt"
	"groupie-tracker/lookup"
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
		if err != nil || intID <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid artist ID")
			return
		}

		artistProfile, err := lookup.FindArtistProfile(intID)
		if errors.Is(err, lookup.ErrArtistNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Artist not found")
			return
		}
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error!")
			return
		}

		Temp, err = template.ParseFiles("templates/artist.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		err = Temp.Execute(w, artistProfile)
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
