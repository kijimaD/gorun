package main

import (
	"os"

	"github.com/kijimaD/gorun/gorun"
)

func main() {
	app := gorun.App{}
	file, err := os.Open("./gorun.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	definition, err := gorun.ParseDefinition(file)
	if err != nil {
		panic(err)
	}
	if err := app.Run(os.Stdin, os.Stdout, os.Stderr, definition); err != nil {
		panic(err)
	}
}
