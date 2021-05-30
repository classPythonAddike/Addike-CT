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

	Files []Resource
}

func ServeResults(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	ReturnData := StartupData{
		Name:        ChallengeInfo.ChallengeName,
		Description: ChallengeInfo.ChallengeDescription,
		Files:       results,
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

	for _, file := range results {
		ReturnData.Values = append(
			ReturnData.Values,
			float32(file.CompletedCases)*float32(100)/float32(file.TotalCases),
		)

		ReturnData.FinishedTesting = append(
			ReturnData.FinishedTesting,
			file.FinishedTesting,
		)

		ReturnData.Passed = append(
			ReturnData.Passed,
			file.Passed,
		)

		if file.FinishedTesting || file.EndTime == file.StartTime {
			ReturnData.Times = append(
				ReturnData.Times,
				float32(file.EndTime.Sub(file.StartTime).Seconds()),
			)
		} else { // file.Endtime will not be defined yet, so its running tests
			ReturnData.Times = append(
				ReturnData.Times,
				float32(time.Since(file.StartTime).Seconds()),
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
