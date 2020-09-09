package download

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// TestBatch tests downloading a batch of images given a set of options.
func TestBatch(t *testing.T) {
	testDir, err := ioutil.TempDir("", "creep")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %s", err)
	}
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Fatalf("Failed to remove temp directory: %s", err)
		}
	}()

	fmt.Println(testDir)

	testCases := []struct {
		expectErr bool
		config    Config
	}{
		{expectErr: false, config: Config{
			Count:    3,
			Name:     "test",
			Out:      testDir,
			Throttle: 0,
			URL:      "https://thispersondoesnotexist.com/image",
		}},
		{expectErr: true, config: Config{
			Count:    1,
			Name:     "test",
			Out:      testDir,
			Throttle: 0,
			URL:      "http://example.com",
		}},
	}

	for _, tc := range testCases {
		errs := Batch(&tc.config)
		if tc.expectErr {
			if errs == nil {
				t.Fatalf("Expected error, got nil")
			}
		} else {
			if errs != nil {
				for _, err := range errs {
					t.Fatal(err.Error())
				}
			}
		}
	}
}

// TestImageFile tests downloading an image.
func TestImageFile(t *testing.T) {
	testDir, err := ioutil.TempDir("", "creep")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %s", err)
	}
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Fatalf("Failed to remove temp directory: %s", err)
		}
	}()

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
		{expectError: false, URL: "http://lorempixel.com/400/200"},
		{expectError: false, URL: "https://thiscatdoesnotexist.com/"},
		{expectError: false, URL: "https://loremflickr.com/320/240"},
		{expectError: false, URL: "https://placeimg.com/640/480/any"},
		{expectError: false, URL: "http://placegoat.com/200"},
		{expectError: false, URL: "https://thisartworkdoesnotexist.com"},
	}
	for i, tc := range testCases {
		p := fmt.Sprintf("test-%d", i)
		f := path.Join(testDir, p)

		err = ImageFile(f, tc.URL)

		if tc.expectError {
			if err == nil {
				t.Fatalf("ImageFile download %s; expected error, got nil.", tc.URL)
			}
		} else {
			if err != nil {
				t.Fatalf("ImageFile returned unexpected error: %s: %v", tc.URL, err)
			}
		}
	}
}
