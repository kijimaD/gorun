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
	tr := TaskRunner{task}
	success := tr.RunTask(renv)
	assert.Equal(t, true, success)
	got := bufout.String()
	expect := `  echo hello
    hello
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
	tr := TaskRunner{task}
	success := tr.RunTask(renv)
	assert.Equal(t, false, success)
	got := buferr.String()
	assert.Contains(t, got, "not_exist_command: ")
}
