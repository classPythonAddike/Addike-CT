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

var ChallengeInfo ChallengeDetails
var FileList []string
var cmd string
var ValidateFiles bool
var TestCases JsonResults
var group sync.WaitGroup
var results []Resource

func InitTests(valid bool) {

	var file []byte
	var err error
	var os string
	var path string
	var code []byte
	var length int
	var res Resource

	ValidateFiles = valid

	file, err = ioutil.ReadFile("config/challenge-info.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &ChallengeInfo)
	if err != nil {
		log.Fatal(err)
	}

	file, err = ioutil.ReadFile("config/cases.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &TestCases)
	if err != nil {
		log.Fatal(err)
	}

	color.Set(color.FgWhite)
	log.Println("Found test cases in file config/cases.json")

	os = runtime.GOOS

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

	for _, i := range GetFiles() {

		code, _ = ioutil.ReadFile(i)
		length = len(code)

		res = Resource{
			File:            i,
			Passed:          false,
			FinishedTesting: false,
			CompletedCases:  0,
			TotalCases:      len(TestCases.Cases),
			CodeLength:      length,
		}
		results = append(results, res)
	}

	group.Add(len(FileList))

	log.Printf("Found %v tests\n", len(FileList))
	color.Unset()
}
