package gorun

import "fmt"

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(j string, renv RuntimeEnvironment) error {
	job := jr.jobs[j]
	fmt.Fprintf(renv.Out, "%s ────────────\n", job.Name)

	for _, task := range job.Steps {
		tr := TaskRunner{task}
		if err := tr.RunTask(renv); err != nil {
			return err
		}
	}
	return nil
}
