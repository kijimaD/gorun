package gorun

import (
	"io"
)

type App struct{}

func (app App) Run(stdin io.Reader, stdout, stderr io.Writer) error {
	def, err := LoadDefinition("gorun.yml")

	if err != nil {
		return err
	}

	jobRunner := JobRunner{
		def.Jobs,
	}

	renv := RuntimeEnvironment{
		In:  stdin,
		Out: stdout,
		Err: stderr,
	}

	for _, j := range def.Jobs {
		jobRunner.RunJob(j.Name, renv)
	}

	return nil
}
