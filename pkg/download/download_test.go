package download

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// TestImageFile tests downloading an image.
func TestImageFile(t *testing.T) {
	testCases := []struct {
		expectError bool
		URL         string
	}{
		{expectError: true, URL: ""},
		{expectError: true, URL: "http://example.com/42"},
		{expectError: true, URL: "http://example.com/"},
		{expectError: false, URL: "https://source.unsplash.com/random"},
		{expectError: false, URL: "https://thispersondoesnotexist.com/image"},
		{expectError: false, URL: "https://picsum.photos/400"},
		// {expectError: false, URL: "http://lorempixel.com/400/200"},
		{expectError: false, URL: "https://thiscatdoesnotexist.com/"},
		{expectError: false, URL: "https://loremflickr.com/320/240"},
		{expectError: false, URL: "https://placeimg.com/640/480/any"},
		{expectError: false, URL: "http://placegoat.com/200"},
	}
	for i, tc := range testCases {
		path := fmt.Sprintf("test-%d", i)
		f, err := ioutil.TempFile("", path)
		if err != nil {
			t.Fatalf("error creating temp file: %s", err)
		}

		err = ImageFile(f.Name(), tc.URL)

		if tc.expectError {
			if err == nil {
				t.Fatalf("ImageFile download %s; expected error, got nil.", tc.URL)
			}
		} else {
			if err != nil {
				t.Fatalf("ImageFile returned unexpected error: %s: %v", tc.URL, err)
			}
		}

		if err := os.Remove(f.Name()); err != nil {
			t.Fatalf("Failed to remove temporary image file %q: %q", f.Name(), err)
		}
	}
}
