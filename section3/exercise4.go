package main

import (
	"encoding/json"
	"net/http"
)

// This is a handler function to receive POSTed form data
func postContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		r.ParseForm()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r.Form)
		w.WriteHeader(200)
	}
}

func main() {
	http.HandleFunc("/v2/contacts", postContact)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err.Error)
	}
}
