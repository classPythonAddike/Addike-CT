package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func HandleCodeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	m, _ := url.ParseQuery(*&r.URL.RawQuery)
	filename := m["file"][0]

	code, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(code)
}
