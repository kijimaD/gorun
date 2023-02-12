package main

import (
	"os"

	"github.com/kijimaD/gorun"
)

func main() {
	app := gorun.App{}
	if err := app.Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		panic(err)
	}
}
