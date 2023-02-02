package Groupie

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	artists := getArtist()
	tmp.Execute(w, artists)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	artist := getArtists(Id)
	// fmt.Println(d)
	tmpl := template.Must(template.ParseFiles("./templates/artists.html"))
	tmpl.Execute(w, artist)
}
