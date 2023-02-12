package gorun

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinition(t *testing.T) {
	got, err := LoadDefinition("gorun.yml")
	if err != nil {
		t.Error(err)
	}
	expect := Definition{Jobs: map[string]Job{
		"a": Job{
			Name:        "a",
			Description: "test",
			Step: Step{
				Task{
					Name:   "a",
					Script: "echo hello"},
			},
		}},
	}
	assert.Equal(t, expect, got)
}

func TestParseDefinition(t *testing.T) {
	r := strings.NewReader(defymlA)
	got, err := ParseDefinition(r)
	if err != nil {
		t.Error(err)
	}
	expect := Definition{Jobs: map[string]Job{
		"a": Job{
			Name:        "a",
			Description: "test",
			Step: Step{
				Task{
					Name:   "a",
					Script: "echo hello"},
			},
		}},
	}
	assert.Equal(t, expect, got)
}

const defymlA = `
jobs:
  a:
    description: test
    step:
      - name: a
        script: echo hello
`
