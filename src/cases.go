package main

import (
	"time"
)

type JsonResults struct {
	Cases []Case
}

type Case struct {
	Input  []string
	Output string
}

type Resource struct {
	File           string // Filename
	TotalCases     int    // Total number of cases to test
	CompletedCases int    // Number of cases that have been completed

	StartTime time.Time // Time at which testing started
	EndTime   time.Time // Time at which testing ended

	Passed          bool // Whether the file passed, or not
	FinishedTesting bool // Whether testing has ended, or not

	CodeLength int
}

type ChallengeDetails struct {
	ChallengeName        string
	ChallengeDescription string

	CleanUpCommand string

	Timeout int
}
