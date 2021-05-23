package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"sync"

	"github.com/fatih/color"
)

var FileList []string
var cmd string
var ValidateFiles bool
var TestCases JsonResults
var group sync.WaitGroup
var results []Resource

func InitTests(valid bool) {

	ValidateFiles = valid

	file, _ := ioutil.ReadFile("cases.json")
	err := json.Unmarshal([]byte(file), &TestCases)

	if err != nil {
		log.Fatal(err)
	}

	color.Set(color.FgWhite)
	log.Println("Found test cases in file cases.json")

	os := runtime.GOOS
	var path string

	switch os {
	case "windows":
		path, err = exec.LookPath("python")
		cmd = "python"

	case "linux":
		path, err = exec.LookPath("python3")
		cmd = "python3"

	case "darwin":
		path, err = exec.LookPath("python3")
		cmd = "python3"
	default:
		Fatal(errors.New("Could not find python in path!"))
	}

	if err != nil {
		Fatal(err)
	}

	color.Set(color.FgWhite)

	log.Println("Found python executable in path at " + path)
	log.Println("Loading test files")

	FileList = GetFiles()

	group.Add(len(FileList))

	log.Printf("Found %v tests\n", len(FileList))
	color.Unset()
}
