package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
		Name:        "Undefined",
		Description: "Undefined",
		Files:       results,
	}

	file, _ := os.Open("challenge.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ChallengeName := ""
	ChallengeDescription := ""

	for scanner.Scan() {
		if ChallengeName == "" {
			ChallengeName += scanner.Text()
		} else {
			ChallengeDescription += scanner.Text()
		}
	}

	if ChallengeName != "" && ChallengeDescription != "" {
		ReturnData.Description = ChallengeDescription
		ReturnData.Name = ChallengeName
	}

	json.NewEncoder(w).Encode(ReturnData)
}

type ProgressResults struct {
	Values []float32 // Percentage of tests completed
	Times  []float32 // Time taken for each test
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
	fmt.Println("Serving UI!")
	http.ServeFile(w, r, "../static/index.html")
}
