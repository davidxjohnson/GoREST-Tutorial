package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// structure to hold unmarshaled json data
type Contact struct {
	RecordKey string   `json:"record_key"`
	LastName  string   `json:"last_name"`
	NickName  []string `json:"nick_name"`
}

// This generates a new record key
func uuid() (uuid string) {
	b := make([]byte, 16)
	rand.Read(b)
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

// This is a handler function to receive POSTed form data
func postContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var contact Contact
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contact.RecordKey = uuid()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func main() {
	http.HandleFunc("/v2/contacts", postContact)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err.Error)
	}
}
