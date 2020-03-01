package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var mimes = map[string]string{
	"image/.jpg": "jpg",
	"image/jpeg": "jpg",
	"image/png":  "png",
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
