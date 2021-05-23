package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func Fatal(err error) {
	color.Set(color.FgRed)
	log.Println(err)
	color.Unset()
	os.Exit(-1)
}
