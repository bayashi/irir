package main

import (
	"testing"
)

/*
	$ go test -fuzz Fuzz -fuzztime=30s
*/
func FuzzProcess(f *testing.F) {
	f.Fuzz(func(t *testing.T, tgt, typ, mch, c string, d []byte) {
		r := &Rule{
			Target: tgt,
			Type:   typ,
			Match:  mch,
			Color:  c,
		}

		process(d, []*Rule{r})
	})
}
