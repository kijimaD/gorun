package gorun

import (
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Step []Task

type Definition struct {
	Jobs map[string]Job
}

type Job struct {
	Name        string
	Description string
	Step        Step
}

type Task struct {
	Name   string
	Script string
}

func LoadDefinition(filename string) (Definition, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return Definition{}, err
	}
	return ParseDefinition(file)
}

func ParseDefinition(r io.Reader) (Definition, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return Definition{}, err
	}

	var raw struct {
		Jobs map[string]struct {
			Description string `yaml:"description"`
			Step        []struct {
				Name   string `yaml:"name"`
				Script string `yaml:"script"`
			} `yaml:step`
		} `yaml:"jobs"`
	}

	if err := yaml.Unmarshal(bs, &raw); err != nil {
		return Definition{}, err
	}

	def := Definition{
		make(map[string]Job, len(raw.Jobs)),
	}

	for name, c := range raw.Jobs {
		tasks := make([]Task, len(c.Step))
		for i, t := range c.Step {
			tasks[i] = Task{
				t.Name,
				t.Script,
			}
		}
		def.Jobs[name] = Job{
			Name:        name,
			Description: c.Description,
			Step:        tasks,
		}
	}

	return def, nil
}
