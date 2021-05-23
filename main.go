package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func main() {

	validate := flag.Bool("validate", false, "Validate each script with `validate.py`")

	flag.Parse()

	InitTests(*validate)

	go TestAllFiles()

	http.HandleFunc("/tests", HandleTestsRequest)
	http.HandleFunc("/results", ServeResults)
	//http.HandleFunc("/", ServeUI)
	http.HandleFunc("/getcode", HandleCodeRequest)

	log.Fatal(http.ListenAndServe(":5000", nil))

	color.Unset()
}
