package gorun

import (
	"bytes"
	"os"
)

type App struct{}

func (app App) Run() {
	task := Task{"hello", "echo hello"}
	job := Job{
		Name:        "job",
		Description: "this is job",
		Steps:       []Task{task},
	}
	def := Definition{map[string]Job{"a": job}}

	jobRunner := JobRunner{
		def.Jobs,
	}

	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}
	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: buferr,
	}

	if err := jobRunner.RunJob("a", renv); err != nil {
		panic(err)
	}
}
