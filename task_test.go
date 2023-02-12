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

	err := RunTask("echo hello", renv)
	if err != nil {
		t.Errorf("%w", err)
	}
	got := bufout.String()
	assert.Equal(t, "hello\n", got)
}

func TestRunTaskFailed(t *testing.T) {
	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: bufout,
		Err: buferr,
	}

	err := RunTask("not_exist", renv)
	if err != nil {
		got := buferr.String()
		assert.Equal(t, "bash: 行 1: not_exist: コマンドが見つかりません\n", got)
	}
	got := bufout.String()
	assert.Equal(t, "", got)
}
