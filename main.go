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

	http.HandleFunc("/startup", ServeResults)
	http.HandleFunc("/progress", ServeProgress)
	//http.HandleFunc("/", ServeUI)
	log.Fatal(http.ListenAndServe(":8080", nil))

	color.Unset()
}
