package main

import (
	"fmt"

	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

func main() {

	version := flag.Bool("version", false, "Output version information")
	flagNoColor := flag.Bool("no-color", false, "Disable colour output")

	flag.Parse()

	if *version {
		fmt.Println("Addike-ct v0.0.1 (alpha)")
		return
	}

	color.NoColor = *flagNoColor

	InitTests()
}
