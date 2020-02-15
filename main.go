package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "https://thispersondoesnotexist.com/image", "The URL of the resource to get")
	filename := flag.String("name", "file", "The filename")
	count := flag.Int("count", 1, "The number of times to get the resource")
	out := flag.String("out", "", "The output directory")
	throttle := flag.Int("throttle", 0, "Duration to wait between downloads")
	flag.Parse()

	if *url == "" {
		exit("a valid URL must be provided")
	}

	if *count <= 0 {
		exit("count must be greater than 0")
	}

	if *out != "" {
		err := parseOut(*out)
		if err != nil {
			exit(err.Error())
		}
	}

	if *throttle < 0 {
		exit("throttle must be a positive integer value")
	}

	var wg sync.WaitGroup
	wg.Add(*count)

	for i := 1; i <= *count; i++ {
		file := fmt.Sprintf("%s-%d.jpg", *filename, i)
		outPath := filepath.Join(*out, file)
		if *throttle > 0 {
			var s string
			if *throttle > 1 {
				s = "seconds"
			} else {
				s = "second"
			}
			fmt.Printf("throttling for %d %s...\n", *throttle, s)
			time.Sleep(time.Second * time.Duration(*throttle))
		}
		fmt.Printf("downloading %s to %s...\n", *url, outPath)
		go func(wg *sync.WaitGroup) {
			err := downloadFile(outPath, *url)
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

func downloadFile(filepath, url string) error {
	// create file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// get data
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check server response
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}

	// write body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// exit prints the given message to the console and terminates the application with an error-code.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

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
