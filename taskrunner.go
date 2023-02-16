package gorun

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

type TaskRunner struct {
	task Task
}

type logScript struct {
	script string
	cmd    *exec.Cmd
	log    *bytes.Buffer
}

func NewScript(script string, env RuntimeEnvironment, out *bytes.Buffer) logScript {
	c := exec.Command("bash", "-c", script)
	c.Stdin = env.In
	c.Stdout = out
	c.Stderr = env.Err
	return logScript{script: script, cmd: c, log: out}
}

func (tr TaskRunner) RunTask(renv RuntimeEnvironment) error {
	out := bytes.Buffer{}
	c := NewScript(tr.task.Run, renv, &out)
	fmt.Fprintln(renv.Out, c.script)

	if err := c.cmd.Start(); err != nil {
		return err
	}

	err := c.cmd.Wait()
	if err != nil {
		return err
	}

	s := bufio.NewScanner(c.log)
	for s.Scan() {
		fmt.Fprintf(renv.Out, "  %s\n", s.Text())
	}

	return nil
}
