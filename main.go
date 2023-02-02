package main

import (
	g "Groupie/Golang"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("./css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", g.HomeHandler)
	http.HandleFunc("/artists", g.Artists)

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe("localhost"+port, nil)

}
