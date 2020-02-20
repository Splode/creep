package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Splode/creep/internal/download"
	"github.com/Splode/creep/internal/flags"
)

func main() {
	config, err := flags.HandleFlags()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse arguments: %s", err.Error()))
	}

	var wg sync.WaitGroup
	wg.Add(int(config.Count))

	for i := 1; i <= int(config.Count); i++ {
		file := fmt.Sprintf("%s-%d", config.Name, i)
		outPath := filepath.Join(config.Out, file)
		if config.Throttle > 0 && i > 1 {
			var s string
			if config.Throttle > 1 {
				s = "seconds"
			} else {
				s = "second"
			}
			fmt.Printf("Throttling for %d %s...\n", config.Throttle, s)
			time.Sleep(time.Second * time.Duration(config.Throttle))
		}
		fmt.Printf("Downloading %s to %s...\n", config.URL, outPath)
		go func(wg *sync.WaitGroup) {
			err := download.Download(outPath, config.URL)
			if err != nil {
				fmt.Printf("Failed to download %s: %s\n", file, err.Error())
			} else {
				fmt.Printf("Successfully downloaded %s to %s\n", file, outPath)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}

// exit prints the given message to the console and terminates the application
// with an error-code.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
