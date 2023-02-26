package gorun

type JobRunner struct {
	jobs map[string]Job
}

func (jr JobRunner) RunJob(jstr string, renv RuntimeEnvironment) bool {
	job := jr.jobs[jstr]

	for i, task := range job.Steps {
		tr := TaskRunner{job.Name, task, len(job.Steps), i + 1}
		tr.RunTask(renv)
	}
	return true
}
