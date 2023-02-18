package gorun

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefinition(t *testing.T) {
	got, err := LoadDefinition("./fixture/gorun.yml")
	if err != nil {
		t.Error(err)
	}
	expect := Definition{Jobs: map[string]Job{
		"job_a": Job{
			Name:        "job_a",
			Description: "test",
			Steps: Steps{
				Task{
					Name: "a",
					Run:  "echo hello1"},
				Task{
					Name: "b",
					Run:  "echo hello2"},
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
		"job_a": Job{
			Name:        "job_a",
			Description: "test",
			Steps: Steps{
				Task{
					Name: "a",
					Run:  "echo hello",
					If:   "which echo",
				},
			},
		}},
	}
	assert.Equal(t, expect, got)
}

const defymlA = `
jobs:
  job_a:
    description: test
    steps:
      - name: a
        run: echo hello
        if: which echo
`
