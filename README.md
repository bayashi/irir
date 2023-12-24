# irir

<a href="https://github.com/bayashi/irir/actions" title="irir CI"><img src="https://github.com/bayashi/irir/workflows/main/badge.svg" alt="irir CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/irir" title="irir report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/irir" alt="irir report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/irir" title="Go irir package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/irir.svg" alt="Go Reference: irir"></a>

`irir` is a command line tool that provides a filter to add colors for text lines generically from a YAML configuration file easily.

## Usage

This is a log file I want to add colors.

```sh
$ cat example.log
2023/05/19 23:56:54 [info] GET /some/resource 200 0.001
2023/05/19 23:56:55 [warn] GET /some/resource 200 0.001
2023/05/19 23:56:56 [error] GET /some/resource 200 0.001
2023/05/19 23:56:57 [info] GET /some/resource 200 0.001
```

then, below is a coloring rule file for `log` in YAML.

```sh
$ cat ~/.config/irir/irir_rule.yaml
---
log:
- type: match
  match: [info]
  color: cyan
  target: word
- type: match
  match: [warn]
  color: yellow
  target: word
- type: match
  match: [error]
  color: bg_red
  target: line
```

Then, you can filter the log file by `irir` with `log` rule.

```sh
$ cat example.log | irir log
```

![colored log file](https://user-images.githubusercontent.com/42190/239714614-fa153eec-a47d-49c8-a5c2-f70dfce97838.png)

Yas!

## Rule YAML

`irir` loads rules from YAML file. The rule file locates on your config directory of [XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html). File name should be `irir_rule.yaml`.

You can start editing `irir_rule.yaml` by a command `irir --edit-config` except on Windows.

You can see the location of `irir_rule.yaml` by a command `irir --dump-config-path`.

Here is the [JSON Schema file](https://raw.githubusercontent.com/bayashi/irir/main/.rule_schema.json) to support writing `irir_rule.yaml`.

### irir_rule.yaml

First key `log` is rule name that is specified in command line. You can name it as you like.

```yaml
log:
- type: match
  match: [info]
  color: cyan
  target: word
- type: match
  match: [warn]
  color: yellow
  target: line
- type: match
  match: [error]
  color: bg_red
  target: line
```

Above rules have 3 ways to color as list.

### Each rule to color

```
- type: match
  match: [info]
  color: cyan
  target: word
```

* `type`: This specifies how to match. It should be `match`, `prefix`, `suffix` or `regexp`. If `target` value is `word`, then you can use only `match` or `regexp`.
* `match`: This is a string or a regexp string to match.
* `color`: specific color name. See [the palette](https://github.com/bayashi/irir/blob/main/color_palette.go)
* `target`: This specifies a scope of coloring. It should be `word` or `line`.

### The case of regexp type

If `type` is `regexp`, then there are special way to set `regexp` for coloring words.

#### Simple regexp case

The condition of regexp is `Ba.`. It will match 2 places with the line `Foo Bar Baz`.

```sh
$ cat example_file.txt
Foo Bar Baz

$ cat ~/.config/irir/irir_rule.yaml
re:
- type: regexp
  match: Ba.
  color: red
  target: word
```

Filter above `example_file.txt` like below:

```sh
$ cat example_file.txt | irir re
```

Output like this.

![colored words by regexp](https://user-images.githubusercontent.com/42190/239843565-1945512c-9e03-49c6-8f4e-7b1b2aad90ba.png)

#### Special chars in YAML for regexp

As for YAML spec, if you write backslash `\` in string value, then you should enclose string value by single-quote like below:

```yaml
- type: regexp
  match: '\w+\.go'
```

If you enclose regexp with backslash by double-quote, then you should escape backslash by backslash:

```yaml
- type: regexp
  match: "\\w+\\.go"
```

This is bit confusing. Single-quoted regexp is easier.

### Another example

To add colors for `go test` result by `gotest` rule.

```yaml
gotest:
- type: prefix
  match: "--- PASS"
  color: green
  target: line
- type: prefix
  match: "ok"
  color: green
  target: line
- type: prefix
  match: "PASS"
  color: green
  target: line
- type: prefix
  match: "--- FAIL"
  color: red
  target: line
- type: prefix
  match: "FAIL"
  color: red
  target: line
- type: prefix
  match: "--- SKIP"
  color: dark_yellow
  target: line
- type: match
  match: "=== RUN"
  color: gray
  target: line
- target: line
  type: match
  match: "=== CONT"
  color: gray
- type: match
  match: "=== PAUSE"
  color: gray
  target: line
- type: match
  match: "[no tests to run]"
  color: yellow
  target: word
- type: match
  match: "[no test files]"
  color: yellow
  target: word
- type: regexp
  match: '[^\/]+\.go:\d+'
  color: cyan
  target: word
```

This is also helpful on wrapped `go test` through `make` in a project.

```sh
$ make test | irir gotest
```

![colored test result](https://user-images.githubusercontent.com/42190/239734820-f18006ce-6a9c-43b8-aaf0-c4f8ebd7a57b.png)

## Default rule

You can specify default rule by ENV:`IRIR_DEFAULT_RULE`. Then you can omit rule argument in command line.

```sh
export IRIR_DEFAULT_RULE=gotest
go test -v ./... | irir
```

## Full options

`irir --help`

```
Usage: cat example.log | irir RULE_ID
Options:
      --dump-colors        Dump color palette
      --dump-config-path   Dump config file path
      --dump-rule          Dump specified rule
      --dump-rules         Show rules from config file
      --dump-schema        Dump JSON Schema to validate the rule YAML config file
      --edit-config        Invoke $EDITOR (or vi) to edit config YAML file
  -h, --help               Show help (This message) and exit
  -v, --version            Show version and build info and exit
```

## Wrap Command Feature (Experimental)

NOTE: Don't execute a command from outside you don't handle. Just invoke only your own commands.

Below command line will color an output from `some-command` by `irir rule`.

```sh
$ iriri rule -- some-command
```

It's as same as below.

```sh
$ some-command | irir rule
```

If you often use `irir`, you can set alias with wrap command feature like below.

```sh
$ alias some-command="iriri rule -- some-command"
```

Then you can avoid writing `| irir rule` on each time.

## TIPS: Color for Github Actions

Github Actions doesn't have TTY. If you want to use `irir` in Github Actions, Then you should add `shell: 'script -q -e -c "bash {0}"'` line like below.

```yaml
  - name: Run tests
    shell: 'script -q -e -c "bash {0}"'
    run: go test -v ./... | irir
```

## Installation

### homebrew install

If you are using Mac:

    brew tap bayashi/tap
    brew install bayashi/tap/irir

### binary install

Download binary from here: https://github.com/bayashi/irir/releases

### go install

If you have golang environment:

```cmd
go install github.com/bayashi/irir@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
