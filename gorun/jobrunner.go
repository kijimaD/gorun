package gorun

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(j string, renv RuntimeEnvironment) bool {
	job := jr.jobs[j]

	for _, task := range job.Steps {
		tr := TaskRunner{job.Name, task}
		tr.RunTask(renv)
	}
	return true
}
