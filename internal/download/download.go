package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	mimes = map[string]string{
		"image/.jpg": "jpg",
		"image/jpeg": "jpg",
		"image/png":  "png",
	}
)

// Download saves the request body from a given URL to the provided
// filepath.
func Download(filepath, url string) error {
	// get data
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	// attempt to get file ext
	ext, err := getExtHeader(res)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// check server response
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}
	// create file
	path := fmt.Sprintf("%s.%s", filepath, ext)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	// write body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
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
