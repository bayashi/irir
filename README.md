# irir

<a href="https://github.com/bayashi/irir/blob/main/LICENSE" title="irir License"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png" alt="MIT License"></a>
<a href="https://github.com/bayashi/irir/actions" title="irir CI"><img src="https://github.com/bayashi/irir/workflows/main/badge.svg" alt="irir CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/irir" title="irir report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/irir" alt="irir report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/irir" title="Go irir package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/irir.svg" alt="Go Reference: irir"></a>

`irir` is a command that provides a filter to add colors for text lines generically from a YAML configuration file easily.

NOTE that irir is still alpha quality. Not test enough. But feedback welcome :-D

## Usage

This is a log file I want to add colors.

```sh
$ cat example.log
2023/05/19 23:56:54 [info] GET /some/resource 200 0.001
2023/05/19 23:56:55 [warn] GET /some/resource 200 0.001
2023/05/19 23:56:56 [error] GET /some/resource 200 0.001
2023/05/19 23:56:57 [info] GET /some/resource 200 0.001
```

then, below is a coloring rule file for `irir` in YAML.

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
  color: red
  target: line
```

Then, you can filter the log file by `irir` with `log` rule.

```sh
$ cat example.log | irir log
```

You can see logs:

![colored log file](https://user-images.githubusercontent.com/42190/239714614-fa153eec-a47d-49c8-a5c2-f70dfce97838.png)

## Rule YAML

`irir` loads rules from YAML file. The rule file locates on your config directory of [XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html). File name should be `irir_rule.yaml`.

### irir_rule.yaml

First key `log` is rule name that is specified in command line:

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

above config is for `log` rule. And it has 3 ways to color as list.

### Each rule to color

```
- type: match
  match: [info]
  color: cyan
  target: word
```

* `type`: `match`, `prefix`, `suffix` or `regexp`. This specifies how to match. If `target` value is `word`, then you can use only `match` or `regexp`.
* `match`: This is a string or a regexp string to match.
* `color`: specific color. See [the palette](https://github.com/bayashi/irir/blob/main/color_palette.go)
* `target`: `word` or `line`. This specifies a scope of coloring.

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

This is helpful on wrapping `go test` by `make`.

```sh
$ make test | irir gotest
```

![colored test result](https://user-images.githubusercontent.com/42190/239734820-f18006ce-6a9c-43b8-aaf0-c4f8ebd7a57b.png)


## Installation

```cmd
go install github.com/bayashi/irir@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
