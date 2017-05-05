package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// This is a handler function to receive POSTed form data
func postContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
		r.ParseMultipartForm(0)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r.MultipartForm)
		file, _, err := r.FormFile("sampledata")
		defer file.Close()
		if err == nil {
			w.Write([]byte("---sample data---\n"))
			io.Copy(w, file)
		}
	}
}

func main() {
	http.HandleFunc("/v2/contacts", postContact)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err.Error)
	}
}
