package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

const (
	exitOK  int = 0
	exitErr int = 1
)

func putErr(message ...interface{}) {
	fmt.Fprintln(os.Stderr, message...)
}

func putUsage() {
	putErr(fmt.Sprintf("Usage: cat example.log | %s RULE_ID", irir))
}

func putHelp(message string) {
	putErr(message)
	putUsage()
	putErr("Options:")
	flag.PrintDefaults()
	os.Exit(exitOK)
}
