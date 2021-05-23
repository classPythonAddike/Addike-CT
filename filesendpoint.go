package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func GetFiles() []string {
	results := []string{}

	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".py" {
			if (ValidateFiles && file.Name() != "validate.py") || !ValidateFiles {
				results = append(results, file.Name())
			}
		}
	}

	return results
}

func GetValidator() string { // Find the validate.py file
	for _, file := range FileList {
		if file == "validate.py" {
			return file
		}
	}

	dir, _ := filepath.Abs(".")

	log.Fatal("Could not find validator script: validate.py in " + dir)
	return ""
}

func HandleTestsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(FileList)
}
