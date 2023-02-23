package gorun

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunTask(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "echo hello", "which make")
	tr := TaskRunner{"job1", 2, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 2/1 echo hello
=> => # hello
`
	assert.Equal(t, expect, got)
}

func TestRunSkip(t *testing.T) {
	bufout := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: &bytes.Buffer{},
	}

	task := newTask("hello", "echo hello", "which not_exist")
	tr := TaskRunner{"job1", 2, task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `=> [job1] 2/2 echo hello
=> => # [skip]
`
	assert.Equal(t, expect, got)
}

func TestRunTaskFailed(t *testing.T) {
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: &bytes.Buffer{},
		Err: buferr,
	}
	task := newTask("hello", "not_exist_command", "")
	tr := TaskRunner{"job1", 1, task}
	success := tr.RunTask(renv)
	assert.Equal(t, false, success)
	got := buferr.String()
	assert.Contains(t, got, "not_exist_command: ")
}
