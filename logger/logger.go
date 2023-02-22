package logger

import (
	"fmt"
	"io"
)

var runlog map[string]infos

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

func Addlog(key string, value info) map[string]infos {
	// TODO: 直に入れたいけどうまくいかない
	result := map[string]infos{}
	for k, v := range runlog {
		result[k] = v
	}
	result[key] = append(result[key], value)
	runlog = result

	return runlog
}

func Output(w io.Writer) {
	fmt.Fprintln(w, runlog)
}
