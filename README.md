# gorun

## install

`$ go install github.com/kijimaD/gorun@main`

or

`docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/gorun:main gorun.yml`

## Usage

prepare `gorun.yml`.

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
      - name: current dir
        run: pwd
        working-directory: /tmp
  job_b:
    description: test
    steps:
      - name: a
        run: echo helloB
      - name: b
        run: echo $WORLD
        env:
          WORLD: hello
```

and run.

```shell
$ gorun gorun.yml
=> [job_a] 7/1 echo helloA
=> => # helloA
=> [job_a] 7/2 go version
=> => # go version go1.20.1 linux/amd64
=> [job_a] 7/3 go build .
=> => # no Go files in /home/silver/Project/gorun
=> [job_a] 7/4 not_exist_command
=> => # bash: 行 1: not_exist_command: コマンドが見つかりません
=> [job_a] 7/5 go fmt ./...
=> [job_a] 7/6 echo this_is_skip
=> => # [skip]
=> [job_a] 7/7 pwd
=> => # /tmp
=> [job_b] 2/1 echo helloB
=> => # helloB
=> [job_b] 2/2 echo $WORLD
=> => # hello
```

inspired: https://github.com/morikuni/ran
