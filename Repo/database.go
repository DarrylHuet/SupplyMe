package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	entity "supplyme/entity"
)

//Env struct provides our interfaces for these assets and their structures in their models.
type Env struct {
	asset interface {
		All_Assets() ([]Cargo.Asset, error)
	}
	items interface {
		All_Units() ([]Cargo.Unit, error)
	}
	users interface {
		user_base() ([]Cargo.User, error)
	}
}

//asset_index provides a user's assets based on their authentication, and provides those resources for the handler
func (env *Env) asset_index(w http.ResponseWriter, r *http.Request) {
	tks, err := env.asset.All_Assets()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, a := range tks {
		fmt.Fprintf(w, "%s, %s, %s, %s, %d\n, %v\n", a.owner, a.aid, a.content, a.id_code, a.timestamp, a.signature)
	}
}

//item_index provides the items based on authentication, and provides all the resources to for the handler
func (env *Env) item_index(w http.ResponseWriter, r *http.Request) {
	item, err := env.items.All_Units()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, u := range item {
		fmt.Fprintf(w, "%s, %s, %s, %s, %s, %d\n. %s", u.owner, u.subtype, u.uid, u.NFT, u.item_code, u.timestamp, u.signature)
	}
}

//user_match finds the user, and provides a session to the client, and access under the API
func (env *Env) user_match(w http.ResponseWriter, r *http.Request) {

	user, err := env.users.user_base()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, u := range user {
		fmt.Fprintf(w, "%s, %s, %s, %s, %s", u.ID, u.Email, u.Password, u.Username, u.PublicKey)
	}
}

//Contribute Session back to repo

func init_db(db_url string) (*Cargo, *Env, error) {
	entity.db, err = sql.Open("postgres", os.Getenv(db_url))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	env := &Env{
		asset: entity.Cargo{DB: db},
		items: entity.Cargo{DB: db},
		users: entity.Cargo{DB: db},
	}

	return entity.db, env, nil
}

var db_url = "postgres://ykkmopruszfdqc:420218a172c4d42409e4611676d1f4fdf068d48ea910d09137174b644ebbeb59@ec2-52-21-153-207.compute-1.amazonaws.com:5432/d1idf7u1b1ckru"

func main() {
	db, err := sql.Open("postgres", os.Getenv(db_url))
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{
		asset: entity.Cargo{DB: db},
		items: entity.Cargo{DB: db},
		users: entity.Cargo{DB: db},
	}

	http.HandleFunc("/assets", env.asset_index)
	http.HandleFunc("/user/{slug}/asset/item", env.item_index)
	http.ListenAndServe(":3000", nil)
	http.HandleFunc("/login", env.user_match)
}