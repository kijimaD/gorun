package gorun

import (
	"testing"
)

// jobキーの配列がよさそうな
func TestOutput(t *testing.T) {
	add("job1", NewInfo("jobA", "taskB", "xxxx-1", "success"))
	add("job2", NewInfo("jobB", "taskA", "xxxx-1", "success"))
	output(runlog)
}
