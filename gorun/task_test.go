package gorun

import (
	"bytes"
	"os"
	"testing"

	"github.com/kijimaD/gorun/logger"
	"github.com/stretchr/testify/assert"
)

func TestRunTask(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "echo hello", "which make", "", map[string]string{})
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 1/1 echo hello
=> => # hello
`
	assert.Equal(t, expect, got)
	logger.Flush()
}

func TestRunSkip(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "echo hello", "which not_exist", "", map[string]string{})
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 1/1 echo hello
=> => # [skip]
`
	assert.Equal(t, expect, got)
	logger.Flush()
}

func TestRunEnv(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "echo $HELLO && echo $WORLD", "", "", map[string]string{"HELLO": "hello", "WORLD": "world"})
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 1/1 echo $HELLO && echo $WORLD
=> => # hello
world
`
	assert.Equal(t, expect, got)
	logger.Flush()
}

func TestWorkdir(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "pwd", "", "/tmp", map[string]string{})
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 1/1 pwd
=> => # /tmp
`
	assert.Equal(t, expect, got)
	logger.Flush()
}

func TestRunTaskFailed(t *testing.T) {
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: &bytes.Buffer{},
		Err: buferr,
	}
	task := newTask("hello", "not_exist_command", "", "", map[string]string{})
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, false, success)
	got := buferr.String()
	assert.Contains(t, got, "not_exist_command: ")
	logger.Flush()
}
