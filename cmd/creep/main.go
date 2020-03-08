package main

import (
	"fmt"
	"os"

	"github.com/Splode/creep/pkg/download"
	"github.com/Splode/creep/pkg/flags"
)

func main() {
	config, err := flags.HandleFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse arguments: %s\n", err)
		os.Exit(1)
	}

	if errs := download.Batch((*download.Config)(config)); errs != nil {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
