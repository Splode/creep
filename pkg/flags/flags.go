package flags

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Version of the app. This is set by goreleaser during release builds using
// the latest git tag.
var Version = "Master"

// Config represents the command-line configuration options.
type Config struct {
	Count    uint
	Name     string
	Out      string
	Throttle uint
	URL      string
}

// HandleFlags parses command line arguments and returns a config.
func HandleFlags() (c *Config, err error) {
	c = &Config{}
	v := false
	flag.StringVar(&c.Name, "name", "creep", "")
	flag.StringVar(&c.Name, "n", "creep", "")
	flag.UintVar(&c.Count, "count", 1, "")
	flag.UintVar(&c.Count, "c", 1, "")
	flag.StringVar(&c.Out, "out", "", "")
	flag.StringVar(&c.Out, "o", "", "")
	flag.UintVar(&c.Throttle, "throttle", 0, "")
	flag.UintVar(&c.Throttle, "t", 0, "")
	flag.BoolVar(&v, "version", false, "")
	flag.BoolVar(&v, "v", false, "")
	flag.Usage = generateUsage()
	flag.Parse()
	c.URL = flag.Arg(0)

	if v {
		fmt.Println(Version)
		os.Exit(0)
	}

	if c.URL == "" {
		err = errors.New("expected a URL, none given")
	}

	if c.Count <= 0 {
		err = fmt.Errorf("expected count to be an integer greater than 0, %d given", c.Count)
	}

	if c.Out != "" {
		err = parseOut(c.Out)
	}

	if err != nil {
		return &Config{}, err
	}

	return c, err
}

func generateUsage() func() {
	return func() {
		fmt.Fprintf(os.Stdout, "\ncreep %s", Version)
		fmt.Fprintf(os.Stdout, `

Downloads an image from the given URL a given number of times to the specified directory.

Usage:
  creep [FLAGS] [OPTIONS] [URL]

URL:
  The URL of the resource to access (required)

Options:
  -c, --count int         The number of times to access the resource (defaults to 1)
  -n, --name string       The base filename to use as output (defaults to "creep")
  -o, --out string        The output directory path (defaults to current directory)
  -t, --throttle int      Number of seconds to wait between downloads (defaults to 0)

Flags:
  -h, --help              Prints help information
  -v, --version           Prints version information

Example usage:
  creep -c 32 https://thispersondoesnotexist.com/image
  creep --name=random --out=downloads --count=64 --throttle=3 https://source.unsplash.com/random`)
		fmt.Println()
		os.Exit(0)
	}
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
