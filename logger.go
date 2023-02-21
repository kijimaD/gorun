package gorun

import (
	"fmt"
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

func add(key string, value info) map[string]infos {
	result := map[string]infos{}
	for k, v := range runlog {
		result[k] = v
	}
	result[key] = append(result[key], value)
	runlog = result

	return runlog
}

func output(log map[string]infos) {
	fmt.Println(log)
}
