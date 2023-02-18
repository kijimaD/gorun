package gorun

import "fmt"

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(j string, renv RuntimeEnvironment) bool {
	job := jr.jobs[j]
	fmt.Fprintf(renv.Out, "%s ────────────\n", job.Name)

	for _, task := range job.Steps {
		tr := TaskRunner{task}
		tr.RunTask(renv)
	}
	return true
}
