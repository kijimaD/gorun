package gorun

import (
	"fmt"
	"os/exec"
)

type TaskRunner struct {
	task Task
}

type logScript struct {
	script string
	cmd    *exec.Cmd
}

func NewScript(script string, env RuntimeEnvironment) logScript {
	c := exec.Command("bash", "-c", script)
	c.Stdin = env.In
	c.Stdout = env.Out
	c.Stderr = env.Err
	return logScript{script: script, cmd: c}
}

func (tr TaskRunner) RunTask(renv RuntimeEnvironment) error {
	c := NewScript(tr.task.Run, renv)
	fmt.Fprintln(renv.Out, c.script)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	err := c.cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
