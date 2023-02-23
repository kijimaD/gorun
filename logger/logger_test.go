package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	w := bytes.Buffer{}

	infoTaskA1 := NewInfo("jobA", "taskA", bytes.NewBufferString("xx-1"), "ok", "echo helloA1")
	infoTaskA1.Addlog().Print(&w)
	infoTaskA2 := NewInfo("jobA", "taskB", bytes.NewBufferString("yy-1"), "ok", "echo helloA2")
	infoTaskA2.Addlog().Print(&w)
	infoTaskB := NewInfo("jobB", "taskB", bytes.NewBufferString("xx-1"), "fail", "echo helloB")
	infoTaskB.Addlog().Print(&w)

	expect := `[jobA] 4/1 echo helloA1
[jobA] 4/2 echo helloA2
[jobB] 4/1 echo helloB
`
	assert.Equal(t, expect, w.String())
}