package main

import "time"

type Test struct {
	FileName  string `json:"File"`
	Extension string

	Language Language
	Compiled bool

	CompletedCases  int
	Passed          bool
	FinishedTesting bool

	CodeLength int
	Time       int

	startTime time.Time
	endTime   time.Time

	timedOut bool
}

type Language struct {
	Language         string
	Command          string
	Execution        string
	Extension        string
	CompiledLanguage bool
}

type LanguageList struct {
	Languages []Language
}

type TestCases struct {
	Cases []Case
}

type Case struct {
	Input  string
	Output string
}

type ChallengeInfo struct {
	Title       string
	Description string
	Timeout     int
}
