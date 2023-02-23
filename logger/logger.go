package logger

import (
	"bytes"
	"fmt"
	"io"
)

var runlog map[string]infos

type infos []info

type info struct {
	job    string
	task   string
	log    *bytes.Buffer
	status string
	script string
}

func NewInfo(job string, task string, log *bytes.Buffer, status string, script string) info {
	info := info{
		job:    job,
		task:   task,
		log:    log,
		status: status,
		script: script,
	}
	return info
}

func Addlog(i info) map[string]infos {
	// TODO: 直に入れたいけどうまくいかない
	result := map[string]infos{}
	for k, v := range runlog {
		result[k] = v
	}
	result[i.job] = append(result[i.job], i)
	runlog = result

	return runlog
}

func PrintTask(w io.Writer, i info) {
	fmt.Fprintf(w, "[%s] 4/%d %s", i.job, 1, i.script)
}
