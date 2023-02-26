package main

import (
	"os"

	"github.com/kijimaD/gorun/gorun"
)

const DEFAULT_CONFIG_PATH = "./gorun.yml"

func main() {
	app := gorun.App{}
	var configfile string
	if len(os.Args) > 1 {
		configfile = os.Args[1]
	} else {
		configfile = DEFAULT_CONFIG_PATH
	}
	file, err := os.Open(configfile)
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
