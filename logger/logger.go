package logger

import (
	"context"
	"fmt"
)

var ctxkey = "Log"

type infos []info

type info struct {
	job    string
	task   string
	log    string
	status string
}

func NewInfo(job string, task string, log string, status string) info {
	info := info{
		job:    job,
		task:   task,
		log:    log,
		status: status,
	}
	return info
}

func add(ctx context.Context, key string, value info) context.Context {
	result := map[string]infos{}
	m := ctx.Value(ctxkey).(map[string]infos)
	for k, v := range m {
		result[k] = v
	}
	result[key] = append(result[key], value)
	ctx = context.WithValue(ctx, ctxkey, result)
	return ctx
}

func output(ctx context.Context) {
	log := ctx.Value(ctxkey).(map[string]infos)
	fmt.Println(log)
}
