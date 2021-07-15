package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type StartupData struct {
	Name        string
	Description string

	Files []Test
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	ReturnData := StartupData{
		Name:        CInfo.Title,
		Description: CInfo.Description,
		Files:       Tests,
	}

	json.NewEncoder(w).Encode(ReturnData)
}

type ProgressResults struct {
	Values          []float32 // Percentage of tests completed
	Times           []float32 // Time taken for each test
	FinishedTesting []bool
	Passed          []bool
}

func ServeProgress(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	ReturnData := ProgressResults{}

	totalCases := len(Cases.Cases)

	for _, file := range Tests {
		ReturnData.Values = append(
			ReturnData.Values,
			float32(file.CompletedCases)*float32(100)/float32(totalCases),
		)

		ReturnData.FinishedTesting = append(
			ReturnData.FinishedTesting,
			file.FinishedTesting,
		)

		ReturnData.Passed = append(
			ReturnData.Passed,
			file.Passed,
		)

		if file.FinishedTesting || file.endTime == file.startTime {
			ReturnData.Times = append(
				ReturnData.Times,
				float32(file.Time),
			)
		} else { // file.Endtime will not be defined yet, so its running tests
			ReturnData.Times = append(
				ReturnData.Times,
				float32(time.Since(file.startTime).Seconds()),
			)
		}
	}

	json.NewEncoder(w).Encode(ReturnData)
}

func ServeUI(w http.ResponseWriter, r *http.Request) {
	var dir string
	var file string
	dir = os.Getenv("CTPath")
	file = (*r).URL.Path
	http.ServeFile(w, r, filepath.Join(dir, file))
}
