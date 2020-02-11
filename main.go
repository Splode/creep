package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	url := flag.String("url", "", "The URL of the resource to get")
	count := flag.Int("count", 1, "The number of times to get the resource")
	out := flag.String("out", "", "The output directory")
	flag.Parse()
	// url := "https://thispersondoesnotexist.com/image"

	if *url == "" {
		panic(errors.New("a valid URL must be provided"))
	}

	if *count <= 0 {
		panic(errors.New("count must be greater than 0"))
	}

	for i := 1; i <= *count; i++ {
		file := fmt.Sprintf("%svisitor-%d.jpg", *out, i)
		time.Sleep(time.Second)
		fmt.Printf("downloading %s to %s\n", *url, file)
		err := downloadFile(file, *url)
		if err != nil {
			fmt.Printf("failed to download %s: %s\n", file, err.Error())
		}
	}
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
