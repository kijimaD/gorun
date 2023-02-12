package gorun

import (
	"bytes"
	"os"
)

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(j string) error {
	job := jr.jobs[j]

	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: buferr,
	}

	for _, task := range job.Step {
		tr := TaskRunner{task}
		if err := tr.RunTask(renv); err != nil {
			return err
		}
	}
	return nil
}
