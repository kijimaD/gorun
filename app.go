package gorun

import (
	"io"
)

type App struct{}

func (app App) Run(stdin io.Reader, stdout, stderr io.Writer) {
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

	renv := RuntimeEnvironment{
		In:  stdin,
		Out: stdout,
		Err: stderr,
	}

	if err := jobRunner.RunJob("a", renv); err != nil {
		panic(err)
	}
}
