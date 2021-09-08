package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

func main() {

	version := flag.BoolP("version", "v", false, "Output version information")
	flagNoColor := flag.BoolP("no-color", "n", false, "Disable colour output")
	create := flag.BoolP("create", "c", false, "Create a new suite of tests, with their configuration")

	flag.Parse()

	if *version {
		fmt.Println("Addike-CT v0.1.0 (stable)")
		return
	}

	if *create {
		CreateNewProject()
		return
	}

	color.NoColor = *flagNoColor

	InitTests()

	http.HandleFunc("/startup", ServeFiles)
	http.HandleFunc("/progress", ServeProgress)
	http.HandleFunc("/", ServeUI)

	PrintLine()
	Info("Running on http://localhost:8080")

	time.Sleep(2 * time.Second)
	go RunTests()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
