package gorun

import (
	"os/exec"
)

// type TaskRunner struct {
// 	Env RuntimeEnvironment
// }

func RunTask(script string, renv RuntimeEnvironment) error {
	cmd := exec.Command("bash", "-c", script)
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
