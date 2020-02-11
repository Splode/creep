package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://thispersondoesnotexist.com/image"
	for i := 1; i <= 2; i++ {
		file := fmt.Sprintf("out/visitor-%d.jpg", i)
		time.Sleep(time.Second)
		fmt.Printf("downloading %s to %s\n", url, file)
		err := downloadFile(file, url)
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
