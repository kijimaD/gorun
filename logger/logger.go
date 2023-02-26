package logger

import (
	"bytes"
	"fmt"
	"io"
)

var runlog map[string]infos

type infos []*info

type info struct {
	job     string
	task    string
	log     *bytes.Buffer
	errlog  *bytes.Buffer
	status  RunStatus
	script  string
	allstep int
	idx     int
}

func Flush() {
	runlog = map[string]infos{}
}

func Result(w io.Writer) {
	const line = "●───────────●"
	fmt.Fprintf(w, "\n%s\nResult\n%s\n\n", line, line)
	for _, v := range runlog {
		for _, info := range v {
			// 最初に入れられたのが入っているだけ。グローバル変数に入っているのを更新しないといけない
			fmt.Fprintf(w, "%s ", info.status)
			info.PrintTask(w)
		}
	}
}

func (i *info) UpdateStatus(r RunStatus) {
	i.status = r
}

func NewInfo(job string, task string, log *bytes.Buffer, errlog *bytes.Buffer, status RunStatus, script string, allstep int, idx int) *info {
	info := info{
		job:     job,
		task:    task,
		log:     log,
		errlog:  errlog,
		status:  status,
		script:  script,
		allstep: allstep,
		idx:     idx,
	}
	return &info
}

func (i *info) Addlog() *info {
	// TODO: 直に入れたいけどうまくいかない
	result := map[string]infos{}
	for k, v := range runlog {
		result[k] = v
	}
	result[i.job] = append(result[i.job], i)
	runlog = result
	return i
}

func (i *info) PrintTask(w io.Writer) *info {
	fmt.Fprintf(w, "=> [%s] %d/%d %s\n", i.job, i.allstep, i.idx, i.script)
	return i
}

func (i *info) PrintCmd(w io.Writer) *info {
	if len(i.log.String()) > 0 {
		fmt.Fprintf(w, "=> => # %s", i.log.String())
	}
	return i
}
func (i *info) PrintCmdErr(w io.Writer) *info {
	if len(i.errlog.String()) > 0 {
		fmt.Fprintf(w, "=> => # %s", i.errlog.String())
	}
	return i
}
