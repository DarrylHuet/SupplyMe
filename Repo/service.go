package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Item struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var items = []Item{
	Item{ID: 1, Name: "Hello", Slug: "010101", Description: "World"},
	Item{ID: 2, Name: "Pfiver Example", Slug: "1010101", Description: "Stop the Spread"},
}

func main() {
	//The creation of a new router that is set up containing our API
	r := mux.NewRouter()

	//Our API consist of N number of routes
	r.Handle("/", http.FileServer(http.Dir("./views")))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/token", ItemHandler).Methods("GET")
	r.Handle("/asset/{slug}/feedback", AddFeedbackHandler).Methods("POST")
	r.Handle("/asset", ItemHandler).Methods("GET")
	r.Handle("/asset/{slug}/feedback", AddFeedbackHandler).Methods("POST")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	http.ListenAndServe(":8080", corsWrapper.Handler(r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var ItemHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//Our payload is the information that is handled throughout the application, and they don't have access to decode them
	payload, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))

})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var item Item
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, i := range items {
		if i.Slug == slug {
			item = i
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if item.Slug != "" {
		payload, _ := json.Marshal(item)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})
