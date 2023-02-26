package gorun

import (
	"io"

	"github.com/kijimaD/gorun/logger"
)

type App struct{}

func (app App) Run(stdin io.Reader, stdout, stderr io.Writer, def Definition) error {
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

	logger.Result(stdout)

	return nil
}
