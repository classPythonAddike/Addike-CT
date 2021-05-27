package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

func validate(file string) (bool, string) {
	command := exec.Command(cmd, "validate.py", file)
	stdout, err := command.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	command.Start()
	dat, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	data := strings.ToLower(strings.TrimSpace(string(dat)))

	command.Wait()

	return data == "valid", data
}

func runtest(file string, input []string, output string) (string, bool) {

	command := exec.Command(cmd, file)
	stdin, err := command.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	stdout, err := command.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	command.Start()

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(input, "\n"))
	}()

	dat, err := ioutil.ReadAll(stdout)
	data := string(dat)

	//if runtime.GOOS == "windows" {
	//	data = strings.ReplaceAll(data, "\r", "")
	//}

	if err != nil {
		log.Fatal(err)
	}

	command.Wait()

	var out string = ""
	if data != output {
		out = data
	}

	return out, data == output

}

func GetHighlightedString(s string) string {

	result := ""

	for _, i := range s {
		if i == '\n' {
			result = result + color.BlueString("\\n")
			color.Unset()
		} else {
			result = result + color.YellowString(string(i))
			color.Unset()
		}
	}

	color.Set(color.FgYellow)

	return result
}

func TestFile(PythonFile string, comp *bool, resultVar *Resource) (bool, string) {

	var out bool
	var result string

	if ValidateFiles {
		out, text := validate(PythonFile)
		if !out {
			color.Set(color.FgRed)
			log.Printf("Invalid file: %v - %v\n", PythonFile, text)
			color.Unset()

			return false, "Invalid file: " + PythonFile + " - " + text
		}
	}

	for _, val := range TestCases.Cases {
		result, out = runtest(PythonFile, val.Input, val.Output)

		(*resultVar).CompletedCases += 1

		if !out && !(*comp) {
			result := fmt.Sprintf(
				"Failed Case for file %v - Expected '%v'%v%v%v%v",
				PythonFile,
				GetHighlightedString(val.Output),
				color.YellowString(", got '"),
				GetHighlightedString(result),
				color.YellowString("'. Input:"),
				color.YellowString(fmt.Sprintf("%v", val.Input)),
			)
			color.Set(color.FgYellow)
			log.Println(result)
			color.Unset()

			return false, result
		}
		if *comp {
			return false, fmt.Sprintf("%v has timed out", PythonFile)
		}
	}

	color.Set(color.FgGreen)
	log.Printf("Passed all cases: %v\n", PythonFile)
	color.Unset()

	return true, ""
}

func TestFileOrTimeout(file string, resultVar *Resource) {
	timer := make(chan string, 1)
	comp := false

	go func() {
		res, _ := TestFile(file, &comp, resultVar)

		if res {
			(*resultVar).CompletedCases = len(TestCases.Cases)
		}

		(*resultVar).Passed = res
		(*resultVar).FinishedTesting = true

		timer <- file + " is done"
	}()

	select {
	case <-timer:
		fmt.Print("")
	case <-time.After(time.Second * 60):
		color.Set(color.FgYellow)
		comp = true
		(*resultVar).FinishedTesting = true
		log.Printf("%v has timed out.\n", file)
		color.Unset()
	}
}

func TestAllFiles() {

	color.Set(color.FgBlue)
	log.Printf("Starting tests")
	color.Unset()

	for pos, file := range FileList {
		results[pos].StartTime = time.Now()
		TestFileOrTimeout(file, &results[pos])
		results[pos].EndTime = time.Now()
		group.Done()
	}
}
