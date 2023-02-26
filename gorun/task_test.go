package gorun

import (
	"bytes"
	"os"
	"testing"

	"github.com/kijimaD/gorun/logger"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		input  Task
		idx    int
		expect string
	}{
		{
			name:  "normal",
			input: newTask("normal", "echo hello", "which make", "", map[string]string{}),
			idx:   1,
			expect: `=> [normal] 4/1 echo hello
=> => # hello
`,
		},
		{
			name:  "skip",
			input: newTask("skip", "echo hello", "which not_exist", "", map[string]string{}),
			idx:   2,
			expect: `=> [skip] 4/2 echo hello
=> => # [skip]
`,
		},
		{
			name:  "env",
			input: newTask("env", "echo $HELLO && echo $WORLD", "", "", map[string]string{"HELLO": "hello", "WORLD": "world"}),
			idx:   3,
			expect: `=> [env] 4/3 echo $HELLO && echo $WORLD
=> => # hello
world
`,
		},
		{
			name:  "workdir",
			input: newTask("workdir", "pwd", "", "/tmp", map[string]string{}),
			idx:   4,
			expect: `=> [workdir] 4/4 pwd
=> => # /tmp
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bufout := &bytes.Buffer{}
			renv := RuntimeEnvironment{
				In:  os.Stdin,
				Out: bufout,
				Err: &bytes.Buffer{},
			}

			tr := TaskRunner{tt.name, tt.input, 4, tt.idx}
			success := tr.RunTask(renv)
			assert.Equal(t, true, success)
			got := bufout.String()
			assert.Equal(t, tt.expect, got)
			logger.Flush()
		})
	}
}

func TestRunTaskFailed(t *testing.T) {
	buferr := &bytes.Buffer{}

	renv := RuntimeEnvironment{
		In:  os.Stdin,
		Out: &bytes.Buffer{},
		Err: buferr,
	}
	task := newTask("hello", "not_exist_command", "", "", map[string]string{"LANG": "en_US"})
	tr := TaskRunner{"job1", task, 1, 1}
	success := tr.RunTask(renv)
	assert.Equal(t, false, success)
	expect := `=> => # bash: line 1: not_exist_command: command not found
`
	got := buferr.String()
	assert.Equal(t, expect, got)
	logger.Flush()
}
