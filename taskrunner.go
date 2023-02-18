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
	errlog *bytes.Buffer
}

func NewScript(script string, env RuntimeEnvironment, out *bytes.Buffer, errlog *bytes.Buffer) logScript {
	c := exec.Command("bash", "-c", script)
	c.Stdin = env.In
	c.Stdout = out
	c.Stderr = errlog
	return logScript{script: script, cmd: c, log: out, errlog: errlog}
}

func (tr TaskRunner) RunTask(renv RuntimeEnvironment) bool {
	success := true
	errbuf := bytes.Buffer{}
	out := bytes.Buffer{}
	c := NewScript(tr.task.Run, renv, &out, &errbuf)
	fmt.Fprintf(renv.Out, "  %s\n", c.script)

	err := c.cmd.Start()
	if err != nil {
		success = false
	}

	err = c.cmd.Wait()
	if err != nil {
		success = false
	}

	s := bufio.NewScanner(c.log)
	for s.Scan() {
		fmt.Fprintf(renv.Out, "    %s\n", s.Text())
	}

	e := bufio.NewScanner(c.errlog)
	for e.Scan() {
		fmt.Fprintf(renv.Err, "    %s\n", e.Text())
	}

	return success
}
