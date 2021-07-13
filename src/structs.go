package main

import "time"

type Test struct {
	FileName  string
	Extension string

	Language Language
	Compiled bool

	CompletedCases  int
	Passed          bool
	FinishedTesting bool

	CodeLength int
	Time       int

	StartTime time.Time
	EndTime   time.Time
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
