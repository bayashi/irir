package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func main() {
	err := run()
	if err != nil {
		putErr(fmt.Sprintf("Err %s: %s", irir, err.Error()))
		os.Exit(exitErr)
	}

	os.Exit(exitOK)
}

func run() error {
	o := parseArgs()

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	rule, err := loadRule(o.rule)
	if err != nil {
		return err
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		result, err := process(s.Bytes(), rule)
		if err != nil {
			return err
		}
		os.Stdout.Write(result)
		os.Stdout.WriteString("\n")
	}

	return nil
}
