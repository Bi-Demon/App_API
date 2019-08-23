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

	// API consist 3 routers

	// /status - call to make sure that API is up and running
	// /products - retrieve a list of products that user can leave feedback on
	// /products/{slug}/feedback - capture user feedback on products

	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/product{slug}/feedback", NotImplemented).Methods("POST")

	// setup server for serving static assest from /static/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// App run on port 3000
	http.ListenAndServe(":3000", r)
}

// Implementing the NotImplemented handler. Whenever an API endpoint is hit
// Simply return message "Not Implemented"

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
