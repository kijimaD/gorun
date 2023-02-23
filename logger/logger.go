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
	errlog *bytes.Buffer
	status string
	script string
}

func NewInfo(job string, task string, log *bytes.Buffer, errlog *bytes.Buffer, status string, script string) info {
	info := info{
		job:    job,
		task:   task,
		log:    log,
		errlog: errlog,
		status: status,
		script: script,
	}
	return info
}

func (i *info) Addlog() *info {
	// TODO: 直に入れたいけどうまくいかない
	result := map[string]infos{}
	for k, v := range runlog {
		result[k] = v
	}
	result[i.job] = append(result[i.job], *i)
	runlog = result
	return i
}

func (i *info) TaskPrint(w io.Writer) *info {
	l := len(runlog[i.job])
	fmt.Fprintf(w, "=> [%s] 4/%d %s\n", i.job, l, i.script)
	return i
}

func (i *info) CmdPrint(w io.Writer) *info {
	if len(i.log.String()) > 0 {
		fmt.Fprintf(w, "=> => # %s", i.log.String())
	}
	return i
}
func (i *info) CmdErrPrint(w io.Writer) *info {
	if len(i.errlog.String()) > 0 {
		fmt.Fprintf(w, "=> => # %s", i.errlog.String())
	}
	return i
}
