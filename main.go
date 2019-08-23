package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// instantiating Gorilla/mux router
	r := mux.NewRouter()

	// simply serve static index page on default page
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// setup server for serving static assest from /static/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// App run on port 3000
	http.ListenAndServe(":3000", r)
}
