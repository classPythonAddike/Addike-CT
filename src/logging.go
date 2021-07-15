package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

func Fatal(message string, err error) {
	color.Set(color.FgRed)
	log.Printf("%v - %v\n", message, err.Error())
	color.Unset()
}

func Warning(message ...string) {
	color.Set(color.FgYellow)
	log.Println(strings.Join(message, " "))
	color.Unset()
}

func Info(message ...string) {
	color.Set(color.FgBlue)
	log.Println(strings.Join(message, " "))
	color.Unset()
}

func Success(message ...string) {
	color.Set(color.FgGreen)
	log.Println(strings.Join(message, " "))
	color.Unset()
}

func PlainInfo(message ...string) {
	color.Unset()
	log.Println(strings.Join(message, " "))
}

func PrintLine() {
	fmt.Println("=====================================================================")
}
