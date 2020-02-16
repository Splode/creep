package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	config, err := handleFlags()
	if err != nil {
		exit(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(int(config.count))

	for i := 1; i <= int(config.count); i++ {
		file := fmt.Sprintf("%s-%d", config.name, i)
		outPath := filepath.Join(config.out, file)
		if config.throttle > 0 && i > 1 {
			var s string
			if config.throttle > 1 {
				s = "seconds"
			} else {
				s = "second"
			}
			fmt.Printf("throttling for %d %s...\n", config.throttle, s)
			time.Sleep(time.Second * time.Duration(config.throttle))
		}
		fmt.Printf("downloading %s to %s...\n", config.url, outPath)
		go func(wg *sync.WaitGroup) {
			err := downloadFile(outPath, config.url)
			if err != nil {
				fmt.Printf("failed to download %s: %s\n", file, err.Error())
			} else {
				fmt.Printf("successfully downloaded %s to %s\n", file, outPath)
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
