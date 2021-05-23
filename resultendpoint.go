package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ServeResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(results)
}

func ServeUI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving UI!")
	http.ServeFile(w, r, "../static/index.html")
}
