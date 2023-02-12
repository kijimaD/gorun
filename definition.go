package gorun

type Step []Task

type Definition struct {
	Jobs map[string]Job
}

type Job struct {
	Name        string
	Description string
	Step        Step
}

type Task struct {
	Name   string
	Script string
}
