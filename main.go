package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	validate := flag.Bool("validate", false, "Validate each script with `validate.py`")

	flag.Parse()

	InitTests(*validate)

	go TestAllFiles()

	http.HandleFunc("/startup", ServeResults)
	http.HandleFunc("/progress", ServeProgress)
	http.HandleFunc("/", ServeUI)

	dir := os.Getenv("CTPath")
	log.Println(dir)

	log.Println("Running on `localhost:8080`")
	log.Fatal(http.ListenAndServe(":8080", nil))

	color.Unset()
}
