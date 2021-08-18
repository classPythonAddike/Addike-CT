package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

func main() {

	version := flag.BoolP("version", "v", false, "Output version information")
	flagNoColor := flag.BoolP("no-color", "n", false, "Disable colour output")
	create := flag.BoolP("create", "c", false, "Create a new suite of tests, with their configuration")

	flag.Parse()

	if *version {
		fmt.Println("Addike-ct v0.0.1 (alpha)")
		return
	}

	if *create {
		CreateNewProject()
		return
	}

	color.NoColor = *flagNoColor

	InitTests()
	go RunTests()

	http.HandleFunc("/startup", ServeFiles)
	http.HandleFunc("/progress", ServeProgress)
	http.HandleFunc("/", ServeUI)

	PrintLine()
	Info("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
