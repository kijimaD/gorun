package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	Addlog("job1", NewInfo("job1", "taskA", "xxxx-1", "success"))
	Addlog("job1", NewInfo("job1", "taskB", "yyyy-1", "fail"))
	Addlog("job2", NewInfo("job2", "taskB", "xxxx-1", "success"))

	w := bytes.Buffer{}
	Output(&w)

	expect := "map[job1:[{job1 taskA xxxx-1 success} {job1 taskB yyyy-1 fail}] job2:[{job2 taskB xxxx-1 success}]]\n"
	assert.Equal(t, expect, w.String())
}
