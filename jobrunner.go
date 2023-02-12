package gorun

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(j string, renv RuntimeEnvironment) error {
	job := jr.jobs[j]

	for _, task := range job.Steps {
		tr := TaskRunner{task}
		if err := tr.RunTask(renv); err != nil {
			return err
		}
	}
	return nil
}
