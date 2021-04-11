package repo

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}

func PassErr() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
