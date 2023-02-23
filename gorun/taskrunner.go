package gorun

import (
	"bytes"
	"os/exec"

	"github.com/kijimaD/gorun/logger"
)

type TaskRunner struct {
	jobName string
	allstep int
	task    Task
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
	execute := true
	out := bytes.Buffer{}
	errbuf := bytes.Buffer{}
	c := NewScript(tr.task.Run, renv, &out, &errbuf)

	info := logger.NewInfo(tr.jobName, tr.task.Name, &out, &errbuf, "aaa", tr.task.Run, tr.allstep)
	info.Addlog().PrintTask(renv.Out)

	// workdir
	if len(tr.task.Workdir) > 0 {
		c.cmd.Dir = tr.task.Workdir
	}

	// process if
	i := NewScript(tr.task.If, renv, &bytes.Buffer{}, &bytes.Buffer{})
	erri := i.cmd.Start()
	const skipmsg = "[skip]\n"
	if erri != nil {
		out = *bytes.NewBufferString(skipmsg)
		execute = false
	}
	erri = i.cmd.Wait()
	if erri != nil {
		out = *bytes.NewBufferString(skipmsg)
		execute = false
	}

	if execute {
		err := c.cmd.Start()
		if err != nil {
			success = false
		}

		err = c.cmd.Wait()
		if err != nil {
			success = false
		}
	}

	info.PrintCmd(renv.Out)
	info.PrintCmdErr(renv.Err)

	return success
}
