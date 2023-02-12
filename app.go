package gorun

type App struct{}

func (app App) Run() {
	task := Task{"hello", "echo hello"}
	job := Job{
		"job",
		"this is job",
		[]Task{task},
	}
	def := Definition{map[string]Job{"a": job}}

	jobRunner := JobRunner{
		def.Jobs,
	}

	if err := jobRunner.RunJob("a"); err != nil {
		panic(err)
	}
}
