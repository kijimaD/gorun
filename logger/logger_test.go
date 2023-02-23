package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	infoTaskA := NewInfo("job1", "taskA", bytes.NewBufferString("xxxx-1"), "success", "echo helloA")
	Addlog(infoTaskA)
	Addlog(NewInfo("job1", "taskB", bytes.NewBufferString("yyyy-1"), "fail", "echo helloB1"))
	Addlog(NewInfo("job2", "taskB", bytes.NewBufferString("xxxx-1"), "success", "echo helloB2"))

	w := bytes.Buffer{}
	PrintTask(&w, infoTaskA)

	expect := "[job1] 4/1 echo helloA"
	assert.Equal(t, expect, w.String())
}
