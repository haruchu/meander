package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/oreilly-japan/go-programming-blueprints/chapter7/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//meander.APIKey = "TODO"

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request){
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
