package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/shlex"
)

func WriteResults() {

	var data []byte
	var file *os.File
	var dir string
	var err error

	dir, err = os.Getwd()

	if err != nil {
		Fatal(err)
	}

	dir = filepath.Join(dir, "results", "challenge-results.json")

	log.Printf("Writing results to '%v'\n", dir)

	data, _ = json.MarshalIndent(results, "", "  ")
	os.Mkdir("results", 0755)
	file, err = os.Create(filepath.Join("results", "challenge-results.json"))

	if err != nil {
		Fatal(err)
	}

	defer file.Close()
	file.WriteString(string(data))

}

func RunCleanUpCommand() {

	var command *exec.Cmd
	var commandArgs []string
	commandArgs, _ = shlex.Split(ChallengeInfo.CleanUpCommand)
	var cmd string = commandArgs[0]

	log.Printf("Running clean up command - %v\n", ChallengeInfo.CleanUpCommand)

	command = exec.Command(cmd, commandArgs[1:]...)
	command.Start()
	command.Wait()

	log.Println("Finished cleaning up")

}
