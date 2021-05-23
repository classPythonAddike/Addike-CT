package main

import "time"

type JsonResults struct {
	Cases []Case
}

type Case struct {
	Input  []string
	Output string
}

type Resource struct {
	File           string
	TotalCases     int
	CompletedCases int

	StartTime time.Time
	EndTime   time.Time

	Passed          bool
	FinishedTesting bool
}
