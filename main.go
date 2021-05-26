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

	http.HandleFunc("/results", ServeResults)
	//http.HandleFunc("/", ServeUI)
	log.Fatal(http.ListenAndServe(":5555", nil))

	color.Unset()
}
