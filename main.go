package main

import (
	"fmt"
	"os"
)

func main() {
	err := run()
	if err != nil {
		putErr(fmt.Sprintf("Err %s: %s", irir, err.Error()))
		os.Exit(exitErr)
	}

	os.Exit(exitOK)
}
