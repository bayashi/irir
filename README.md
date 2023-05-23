# irir

<a href="https://github.com/bayashi/irir/blob/main/LICENSE" title="irir License"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png" alt="MIT License"></a>
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

#### Regexp replacement by matched elements

The condition of regexp is `(Ba)(.)\nGa$2`. It will be split into a matching regexp `(Ba)(.)` and replacing string `Ga$2` by `\n`.

```sh
$ cat example_file.txt
Foo Bar Baz

$ cat ~/.config/irir/irir_rule.yaml
re:
- type: regexp
  match: "(Ba)(.)\nGa$2"
  color: red
  target: word
```

Output will be:

![colored and replaced words by regexp](https://user-images.githubusercontent.com/42190/239849754-b67e4fbd-8616-4149-8723-e5aa8c8605e4.png)

### Another example

To add colors for `go test` result.

`irir_rule.yaml` is like below.

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
```

This is helpful on wrapped `go test` through `make`.

```sh
$ make test | irir gotest
```

![colored test result](https://user-images.githubusercontent.com/42190/239734820-f18006ce-6a9c-43b8-aaf0-c4f8ebd7a57b.png)

## Full options

`irir --help`

```
Usage: cat example.log | irir RULE_ID
Options:
      --dump-colors        Dump color palette for enum list
      --dump-config-path   Dump config file path
      --dump-rule          Dump specified rule
      --dump-rules         Show rules from config file
      --dump-schema        Dump JSON Schema to validate the rule YAML config file
  -h, --help               Show help (This message) and exit
  -v, --version            Show version and build info and exit
```

## Installation

```cmd
go install github.com/bayashi/irir@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
