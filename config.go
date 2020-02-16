package main

// config represents the command-line configuration options.
type config struct {
	count    uint
	name     string
	out      string
	throttle uint
	url      string
}
