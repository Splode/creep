package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const version = "0.1.1"

func generateUsage() func() {
	return func() {
		fmt.Printf("\ncreep %s", version)
		fmt.Println(`

Downloads an image from the given URL a given number of times to the specified directory.

Usage:
  creep [FLAGS] [OPTIONS]

Options:
  -u, --url string        The URL of the resource to access (required)
  -c, --count int         The number of times to access the resource (defaults to 1)
  -n, --name string       The base filename to use as output (defaults to "creep")
  -o, --out string        The output directory path (defaults to current directory)
  -t, --throttle int      Number of seconds to wait between downloads (defaults to 0)

Flags:
  -h, --help              Prints help information
  -v, --version           Prints version information

Example usage:
  creep -u https://thispersondoesnotexist.com/image -c 32
  creep --url=https://source.unsplash.com/random --name=random --out=downloads --count=64 --throttle=3
		`)
	}
}

// handleFlags parses command line arguments and returns a config.
func handleFlags() (c *config, err error) {
	c = &config{}
	v := false
	flag.StringVar(&c.url, "url", "", "")
	flag.StringVar(&c.url, "u", "", "")
	flag.StringVar(&c.name, "name", "creep", "")
	flag.StringVar(&c.name, "n", "creep", "")
	flag.UintVar(&c.count, "count", 1, "")
	flag.UintVar(&c.count, "c", 1, "")
	flag.StringVar(&c.out, "out", "", "")
	flag.StringVar(&c.out, "o", "", "")
	flag.UintVar(&c.throttle, "throttle", 0, "")
	flag.UintVar(&c.throttle, "t", 0, "")
	flag.BoolVar(&v, "version", false, "")
	flag.BoolVar(&v, "v", false, "")
	flag.Usage = generateUsage()
	flag.Parse()

	if v {
		fmt.Println(version)
		os.Exit(0)
	}

	if c.url == "" {
		err = errors.New("expected a URL, none given")
	}

	if c.count <= 0 {
		err = fmt.Errorf("expected count to be an integer greater than 0, %d given", c.count)
	}

	if c.out != "" {
		err = parseOut(c.out)
	}

	if err != nil {
		return &config{}, err
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
