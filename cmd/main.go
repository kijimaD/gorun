package main

import (
	"os"

	"github.com/kijimaD/gorun"
)

func main() {
	app := gorun.App{}
	app.Run(os.Stdin, os.Stdout, os.Stderr)
}
