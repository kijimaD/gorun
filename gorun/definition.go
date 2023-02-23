package gorun

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Steps []Task

type Definition struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type Job struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       Steps
}

type Task struct {
	Name    string `yaml:"name"`
	Run     string `yaml:"run"`
	If      string `yaml:"if"`
	Workdir string `yaml:"working-directory"`
}

func newTask(name string, run string, ifarg string, workdir string) Task {
	return Task{
		name,
		run,
		ifarg,
		workdir,
	}
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
		tasks := make([]Task, len(c.Steps))
		for i, t := range c.Steps {
			tasks[i] = newTask(
				t.Name,
				t.Run,
				t.If,
				t.Workdir,
			)
		}
		def.Jobs[name] = Job{
			Name:        name,
			Description: c.Description,
			Steps:       tasks,
		}
	}

	return def, nil
}
