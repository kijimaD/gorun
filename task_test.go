package gorun

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunTask(t *testing.T) {
	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: buferr,
	}

	task := Task{"hello", "echo hello"}
	tr := TaskRunner{task}
	err := tr.RunTask(renv)
	if err != nil {
		t.Error(err)
	}
	got := bufout.String()
	expect := `  echo hello
    hello
`
	assert.Equal(t, expect, got)
}

func TestRunTaskFailed(t *testing.T) {
	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: buferr,
	}

	task := Task{"hello", "not_exist_command"}
	tr := TaskRunner{task}
	err := tr.RunTask(renv)
	if err != nil {
		got := buferr.String()
		assert.Contains(t, got, "not_exist_command: ")
	}
	expect := `  not_exist_command
`
	got := bufout.String()
	assert.Equal(t, expect, got)
}
