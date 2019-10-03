package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var saved = &Request{}

type Response struct {
	Data interface{} `json:"data"`
}

type Request struct {
	Name string `json:"name"`
}

func paulHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&Response{Data: saved})
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		body := &Request{}
		err := decoder.Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Bad Request")
			return
		}

		saved = body
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Success")
	default:
		return
	}

}

func main() {
	saved.Name = "initial"
	http.HandleFunc("/mytoys/name", paulHandler)
	log.Println("running at http://localhost:11000")
	log.Fatal(http.ListenAndServe(":11000", nil))
}
