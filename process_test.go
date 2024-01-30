package main

import (
	"regexp"
	"testing"

	a "github.com/bayashi/actually"
	"github.com/fatih/color"
)

func Test_process(t *testing.T) {
	type args struct {
		origLine []byte
		rule     []*Rule
	}
	tests := []struct {
		name   string
		args   args
		expect []byte
	}{
		{
			name: "word",
			args: args{
				origLine: []byte("Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "word",
						Type:   "match",
						Match:  "Bar",
						Color:  "red",
					},
				},
			},
			expect: []byte("Foo \x1b[91mBar\x1b[0m Baz"),
		},
		{
			name: "word of a part",
			args: args{
				origLine: []byte("Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "word",
						Type:   "match",
						Match:  "Ba",
						Color:  "red",
					},
				},
			},
			expect: []byte("Foo \x1b[91mBa\x1b[0mr \x1b[91mBa\x1b[0mz"),
		},
		{
			name: "line",
			args: args{
				origLine: []byte("Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "line",
						Type:   "match",
						Match:  "Bar",
						Color:  "red",
					},
				},
			},
			expect: []byte("\x1b[91mFoo Bar Baz\x1b[0m"),
		},
		{
			name: "prefix",
			args: args{
				origLine: []byte("    Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "line",
						Type:   "prefix",
						Match:  "Foo",
						Color:  "red",
					},
				},
			},
			expect: []byte("\x1b[91m    Foo Bar Baz\x1b[0m"),
		},
		{
			name: "suffix",
			args: args{
				origLine: []byte("Foo Bar Baz    "),
				rule: []*Rule{
					&Rule{
						Target: "line",
						Type:   "suffix",
						Match:  "Baz",
						Color:  "red",
					},
				},
			},
			expect: []byte("\x1b[91mFoo Bar Baz    \x1b[0m"),
		},
		{
			name: "regexp line",
			args: args{
				origLine: []byte("Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "line",
						Type:   "regexp",
						Match:  "F..",
						Color:  "red",
						regexp: regexp.MustCompile("(F..)"),
					},
				},
			},
			expect: []byte("\x1b[91mFoo Bar Baz\x1b[0m"),
		},
		{
			name: "regexp word",
			args: args{
				origLine: []byte("Foo Bar Baz"),
				rule: []*Rule{
					&Rule{
						Target: "word",
						Type:   "regexp",
						Match:  "Ba.",
						Color:  "red",
						regexp: regexp.MustCompile("(Ba.)"),
					},
				},
			},
			expect: []byte("Foo \x1b[91mBar\x1b[0m \x1b[91mBaz\x1b[0m"),
		},
		{
			name: "color text including `%`",
			args: args{
				origLine: []byte("coverage: 94.6% of statements"),
				rule: []*Rule{
					&Rule{
						Target: "line",
						Type:   "match",
						Match:  "coverage",
						Color:  "red",
					},
				},
			},
			expect: []byte("\x1b[91mcoverage: 94.6% of statements\x1b[0m"),
		},
	}

	color.NoColor = false

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := process(tt.args.origLine, tt.args.rule)
			a.Got(err).NoError(t)
			a.Got(g).Expect(tt.expect).Same(t)
		})
	}
}
