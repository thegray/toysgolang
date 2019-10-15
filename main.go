package main

import (
	"encoding/json"
	"log"
	"net/http"
	"toysgolang/rds"

	"github.com/gomodule/redigo/redis"
)

var saved = &Request{}
var rc *redis.Pool

type Response struct {
	Data interface{} `json:"data"`
}

type Request struct {
	Name string `json:"name"`
}

func paulHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conn := rc.Get()
	defer conn.Close()
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(200)
		data, err := rds.Get(conn, "toysaved")
		if err != nil {
			log.Println("key not found")
			json.NewEncoder(w).Encode(&Response{Data: saved})
			return
		}
		saved.Name = data
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
		err = rds.Set(conn, "toysaved", saved.Name)
		if err != nil {
			log.Println("failed setting key value")
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Success")
	default:
		return
	}

}

func main() {
	rc = rds.NewPool()
	if rc == nil {
		log.Fatalln("redis conn is nil!")
	}

	saved.Name = "initial"
	http.HandleFunc("/name", paulHandler)
	log.Println("running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
