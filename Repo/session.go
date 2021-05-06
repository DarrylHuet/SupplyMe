package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"encoding/json"

	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/heroku"
)

var (
	store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")), []byte(os.Getenv("COOKIE_ENCRYPT")))

	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("QSp8DZNKE0M6jvoNBaW1oX79k1NiwvXM"),
		ClientSecret: os.Getenv("SECRET"),
		Endpoint:     heroku.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "user"},
		RedirectURL:  "http://" + os.Getenv("app-supplyme.") + "herouapp.com/auth/heroku/callback",
	}
	stateToken = os.Getenv("app-supplyme")
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauthConfig
	ctx      context.Context
}

func init() error {
	gob.Register(&oauth2.Token{})

	store.MaxAge(60 * 60 * 8)
	store.Options.Secure = true

	return nil
}

var handleRoot = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><body><a href="/auth/heroku">Sign in with Heroku</a></body></html>`)
})
var handleAuth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(stateToken)
	http.Redirect(w, r, url, http.StatusFound)
})

var handleAuthCallback = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if v := r.FormValue("state"); v != stateToken {
		http.Error(w, "Invalid State token", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	token, err := oauthConfig.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := store.Get(r, "spring-bush-3329.us.auth0.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["heroku-oauth-token"] = token
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/user", http.StatusFound)
})

var handleUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "spring-bush-3329.us.auth0.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, ok := session.Values["heroku-oauth-token"].(*oauth2.Token)
	if !ok {
		http.Error(w, "Unable to asset webtoken", http.StatusInternalServerError)
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.heroku.com/account")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	var User struct {
		Email string "json:email"
	}
	if err := d.Decode(&User); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, `<html><body><h1>Hello %s</h1></body></html>`, User.Email)
})
