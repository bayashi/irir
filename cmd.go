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

func putHelp() {
	putErr(fmt.Sprintf("Version %s", getVersion()))
	putUsage()
	putErr("Options:")
	flag.PrintDefaults()
	os.Exit(exitOK)
}

func putHelpWithMessage(message string) {
	if message != "" {
		putErr(message)
	}
	putHelp()
}
