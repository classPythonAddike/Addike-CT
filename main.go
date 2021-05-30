package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func main() {

	validate := flag.Bool("validate", false, "Validate each script with `validate.py`")
	version := flag.Bool("version", false, "Output version information")

	flag.Parse()

	if *version {
		fmt.Println("Addike-ct v0.0.1 (alpha)")
		return
	}

	InitTests(*validate)

	go TestAllFiles()

	http.HandleFunc("/startup", ServeResults)
	http.HandleFunc("/progress", ServeProgress)
	http.HandleFunc("/", ServeUI)

	log.Println("Running on `localhost:8080`")
	log.Fatal(http.ListenAndServe(":8080", nil))

	color.Unset()
}
