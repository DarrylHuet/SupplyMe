package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Env struct {
	asset interface {
		All_Assets() ([]Asset, error)
	}
	items interface {
		All_Units() ([]Unit, error)
	}
}

func main() {
	db, err := sql.Open("postgres", "Postgres://")
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{
		asset: Cargo{DB: db},
	}

	http.HandleFunc("/assets", env.asset_index)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) asset_index(w http.ResponseWriter, r *http.Request) {
	tks, err := env.asset.All_Assets()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, a := range tks {
		fmt.Fprintf(w, "%s, \n, %s, %s, %s", a.owner, a.aid, a.content, a.timestamp, a.signature)
	}
}

func (env *Env) item_index(w http.ResponseWriter, r *http.Request) {
	item, err := env.items.All_Units()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, u := range item {
		fmt.Fprintf(w, "%s, \n, %s, \n. %s", u.owner, u.uid, u.NFT, u.creation_time, u.timestamp, u.signature)
	}
}
