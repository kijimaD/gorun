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
	file, err := os.Open("../gorun.yml")
	if err != nil {
		t.Error()
	}
	definition, err := ParseDefinition(file)
	if err != nil {
		t.Error(err)
	}
	if err := app.Run(os.Stdin, bufout, buferr, definition); err != nil {
		t.Error(err)
	}
}
