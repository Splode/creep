package main

import (
	"errors"
	"flag"
	"os"
)

// handleFlags parses command line arguments and returns a config.
func handleFlags() (c *config, err error) {
	c = &config{}
	flag.StringVar(&c.url, "url", "", "The URL of the resource to get")
	flag.StringVar(&c.name, "name", "creep", "The filename")
	flag.UintVar(&c.count, "count", 1, "The number of times to get the resource")
	flag.StringVar(&c.out, "out", "", "The output directory")
	flag.UintVar(&c.throttle, "throttle", 0, "Duration to wait between downloads")
	flag.Parse()

	if c.url == "" {
		err = errors.New("A URL must be provided")
	}

	if c.count <= 0 {
		err = errors.New("Count must be greater than 0")
	}

	if c.out != "" {
		err = parseOut(c.out)
	}

	if c.throttle < 0 {
		err = errors.New("Throttle must be a positive integer")
	}

	return c, err
}

// parseOut validates the given directory path, creating the directory at the
// given path if it does not exist.
func parseOut(out string) error {
	if _, err := os.Stat(out); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(out, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
