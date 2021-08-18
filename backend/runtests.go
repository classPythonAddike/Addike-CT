package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/shlex"
)

var group sync.WaitGroup

func TestFile(test *Test) bool {
	for pos, testcase := range Cases.Cases {

		if test.timedOut {
			return false
		}

		if !test.Compiled && test.Language.CompiledLanguage {
			return false
		}

		execCommand := strings.ReplaceAll(test.Language.Execution, "[f]", test.FileName)
		execCommand = strings.ReplaceAll(execCommand, "[e]", test.Extension)
		execCommand = strings.ReplaceAll(
			execCommand,
			"[n]",
			test.FileName[0:len(test.FileName)-len(test.Extension)],
		)

		cmd, _ := shlex.Split(execCommand)
		var executor *exec.Cmd

		if len(cmd) == 1 {
			executor = exec.Command(cmd[0])
		} else {
			executor = exec.Command(cmd[0], cmd[1:]...)
		}

		stdin, err := executor.StdinPipe()

		if err != nil {
			Fatal("", err)
		}

		stdout, err := executor.StdoutPipe()

		if err != nil {
			Fatal("", err)
		}

		executor.Start()

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, testcase.Input)
		}()

		rawData, err := ioutil.ReadAll(stdout)
		data := string(rawData)

		if err != nil {
			Fatal("", err)
		}

		executor.Wait()

		if data != testcase.Output {
			test.CompletedCases = pos
			return false
		}

		test.CompletedCases += 1
	}

	return true
}

func TestFileOrTimeout(test *Test) {
	timer := make(chan string, 1)

	go func() {
		res := TestFile(test)

		if res {
			test.CompletedCases = len(Cases.Cases)
		}

		test.Passed = res
		test.FinishedTesting = true

		timer <- test.FileName + " is done"
	}()

	select {

	case <-timer:

		if test.Passed {
			Success("Passed all testcases -", test.FileName)
		} else {
			Warning(test.FileName, "failed tests")
		}

	case <-time.After(time.Second * time.Duration(CInfo.Timeout)):
		test.timedOut = true
		Warning(test.FileName, "timed out")
	}

	group.Add(-1)
}

func RunTests() {
	group.Add(len(Tests))

	PrintLine()
	PlainInfo("Starting tests")

	for pos := range Tests {

		Info("Testing", Tests[pos].FileName)

		Tests[pos].startTime = time.Now()

		TestFileOrTimeout(&Tests[pos])

		Tests[pos].endTime = time.Now()

		duration := Tests[pos].endTime.Sub(Tests[pos].startTime).Seconds()
		Tests[pos].Time = int(duration)
	}

	PrintLine()
	PlainInfo("Writing results to results/challenge-results.json")

	data, _ := json.MarshalIndent(Tests, "", "  ")
	os.Mkdir("results", 0755)
	file, err := os.Create(filepath.Join("results", "challenge-results.json"))

	if err != nil {
		Fatal("", err)
	}

	defer file.Close()
	file.Write(data)
}
