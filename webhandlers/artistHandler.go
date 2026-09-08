package webhandlers

import (
	"fmt"
	"groupie-tracker/apihandlers"
	"html/template"
	"log"
	"net/http"
)

func Artists(web http.ResponseWriter, request *http.Request) {
	var Temp *template.Template

	if request.Method == http.MethodGet {
		
		artists, err := handlers.FetchArtists()
		if err != nil {
			web.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(web, err.Error())
			return
		}

		Temp, err = template.ParseFiles("templates/artists.html")
		if err != nil {
			web.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(web, err.Error())
			return
		}

		err = Temp.Execute(web, artists)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		web.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(web, "Method not allowed")
		return
	}
}
