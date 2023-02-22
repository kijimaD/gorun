package gorun

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// jobキーの配列がよさそうな
func TestOutput(t *testing.T) {
	addlog("job1", NewInfo("job1", "taskA", "xxxx-1", "success"))
	addlog("job1", NewInfo("job1", "taskB", "yyyy-1", "fail"))
	addlog("job2", NewInfo("job2", "taskB", "xxxx-1", "success"))

	w := bytes.Buffer{}
	output(&w, runlog)

	expect := "map[job1:[{job1 taskA xxxx-1 success} {job1 taskB yyyy-1 fail}] job2:[{job2 taskB xxxx-1 success}]]\n"
	assert.Equal(t, expect, w.String())
}
