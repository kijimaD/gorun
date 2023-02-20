package logger

import (
	"context"
	"testing"
)

// jobキーの配列がよさそうな
func TestOutput(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxkey, map[string]infos{})
	ctx = add(ctx, "job1", NewInfo("jobA", "taskA", "xxxx", "success"))
	ctx = add(ctx, "job1", NewInfo("jobA", "taskB", "xxxx-1", "success"))
	ctx = add(ctx, "job2", NewInfo("jobB", "taskB", "yyyy", "failed"))

	output(ctx)
}
