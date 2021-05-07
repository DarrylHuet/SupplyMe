package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"os"
)

type Page struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var pages = []Page{
	Page{ID: 1, Name: "Hello", Slug: "010101", Description: "World"},
	Page{ID: 2, Name: "Pfiver Example", Slug: "1010101", Description: "Stop the Spread"},
}

func main() {
	//The creation of a new router that is set up containing our API
	r := mux.NewRouter()
	db, err = init_db()
	if err != nil {
		log.Println(err)
	}

	//Our API consist of N number of routes
	r.Handle("/", handleRoot).Methods("POST")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/page", PageHandler).Methods("GET")
	r.Handle("/asset/{slug}/feedback", AddFeedbackHandler).Methods("POST")
	r.Handle("/asset", PageHandler).Methods("GET")
	r.Handle("/asset/{slug}/feedback", AddFeedbackHandler).Methods("POST")
	r.Handle("/auth/heroku", handleAuth).Methods("GET")
	r.Handle("/auth/heroku/callback", handleAuthCallback).Methods("POST")
	r.Handle("/user", handleUser).Methods("GET")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	http.HandleFunc("/assets", db.asset_index)
	http.HandleFunc("/user/{slug}/asset/item", db.item_index)
	http.ListenAndServe(":3000", nil)
	http.HandleFunc("/login", db.user_match)

	http.ListenAndServe(":"+os.Getenv("8080"), corsWrapper.Handler(r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, err := db.asset_index
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))

})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var page Page
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range pages {
		if p.Slug == slug {
			page = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if page.Slug != "" {
		payload, _ := json.Marshal(page)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})
