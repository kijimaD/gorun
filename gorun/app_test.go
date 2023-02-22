package gorun

import (
	"bytes"
	"os"
	"testing"
)

func TestApp(t *testing.T) {
	app := App{}

	bufout := &bytes.Buffer{}
	buferr := &bytes.Buffer{}
	if err := app.Run(os.Stdin, bufout, buferr); err != nil {
		t.Error(err)
	}
}
