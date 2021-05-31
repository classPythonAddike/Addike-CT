package main

import (
	"io/ioutil"
	"log"
	"os"
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
			results = append(results, file.Name())
		}
	}

	return results
}

func GetValidator() string { // Find the validate.py file

	f, err := os.Open("./config/validate.py")

	if err != nil {
		dir, _ := filepath.Abs(".")
		log.Fatal("Could not find validator script: validate.py in " + dir + "/config")
		return ""
	} else {
		f.Close()
		return "./config/validate.py"
	}
}
