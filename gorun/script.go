package gorun

import "io"

type RuntimeEnvironment struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}
