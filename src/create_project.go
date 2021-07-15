package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var githubRepo string = "https://raw.githubusercontent.com/classPythonAddike/challenge-tester-backend/master/src/sampletests"

var files []string = []string{
	"/config/cases.json",
	"/config/challenge-info.json",
	"/config/generator.py",
	"/config/language-config.json",
	"/config/validate.py",
}

func CreateNewProject() {
	cwd, _ := filepath.Abs(".")
	log.Printf("Creating new project at %v\n", cwd)

	CreateFolders()

	for _, route := range files {

		resp, err := http.Get(githubRepo + route)

		if err != nil {
			Warning("Error while fetching", route, "-", err.Error())
			continue
		}

		PlainInfo("Creating file", route)

		content, _ := ioutil.ReadAll(resp.Body)
		os.WriteFile("."+route, content, 0755)
	}
}

func CreateFolders() {
	os.Mkdir("config", 0755)
}
