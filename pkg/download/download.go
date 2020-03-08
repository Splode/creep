package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var mimes = map[string]string{
	"image/.jpg": "jpg",
	"image/jpeg": "jpg",
	"image/jpg":  "jpg",
	"image/png":  "png",
}

// Config represents the command-line configuration options.
type Config struct {
	Count    uint
	Name     string
	Out      string
	Throttle uint
	URL      string
}

// Batch downloads a batch of images given a set of options.
func Batch(config *Config) (errs []error) {
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
			if err := ImageFile(outPath, config.URL); err != nil {
				errs = append(errs, fmt.Errorf("Failed to download %s: %s\n", file, err.Error()))
			} else {
				fmt.Printf("Successfully downloaded %s to %s\n", file, outPath)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	return errs
}

// ImageFile saves the request body from a given URL to the provided filepath.
func ImageFile(filepath, url string) (err error) {
	// get data
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			err = err.(error)
		}
	}()

	// check server response
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}

	// attempt to get file ext
	ext, err := getExtHeader(res)
	if err != nil {
		return err
	}

	// create file
	path := fmt.Sprintf("%s.%s", filepath, ext)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			err = err.(error)
		}
	}()

	// write body to file
	if _, err = io.Copy(out, res.Body); err != nil {
		return err
	}

	return nil
}

// getExtHeader attempts to infer a file extension from the Content-Type of a
// given response header using mime types.
func getExtHeader(r *http.Response) (string, error) {
	ct := r.Header["Content-Type"][0]
	mime, prs := mimes[ct]
	if !prs {
		return "", errors.New("could not detect mime-type from http response")
	}
	return mime, nil
}
