# gorun

prepare yaml.

```yaml
jobs:
  job_a:
    description: test
    steps:
      - name: a
        run: echo helloA
      - name: run version
        run: go version
      - name: run version
        run: go build .
      - name: error
        run: not_exist_command
      - name: help
        run: go fmt ./...
        if: which make
      - name: skip
        run: echo this_is_skip
        if: which skip
  job_b:
    description: test
    steps:
      - name: a
        run: echo helloB

```

and run.

```shell
$ go run ./cmd
=> [job_a] 6/1 echo helloA
=> => # helloA
=> [job_a] 6/2 go version
=> => # go version go1.20.1 linux/amd64
=> [job_a] 6/3 go build .
=> => # no Go files in /home/silver/Project/gorun
=> [job_a] 6/4 not_exist_command
=> => # bash: 行 1: not_exist_command: コマンドが見つかりません
=> [job_a] 6/5 go fmt ./...
=> [job_a] 6/6 echo this_is_skip
=> => # [skip]
=> [job_b] 1/1 echo helloB
=> => # helloB
```

inspired: https://github.com/morikuni/ran
