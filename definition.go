package gorun

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Step []Task

type Definition struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Step        Step
}

type Task struct {
	Name   string `yaml:"name"`
	Script string `yaml:"script"`
}

func LoadDefinition(filename string) (Definition, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Definition{}, err
	}
	defer file.Close()
	return ParseDefinition(file)
}

func ParseDefinition(r io.Reader) (Definition, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return Definition{}, err
	}

	var raw Definition

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
