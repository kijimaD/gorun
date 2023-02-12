package gorun

import "os/exec"

type TaskRunner struct {
	task Task
}

func (tr TaskRunner) RunTask(renv RuntimeEnvironment) error {
	cmd := exec.Command("bash", "-c", tr.task.Run)
	cmd.Stdin = renv.In
	cmd.Stdout = renv.Out
	cmd.Stderr = renv.Err

	if err := cmd.Start(); err != nil {
		return err
	}

	err := cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
