package download

import (
	"fmt"
	"testing"
)

func TestImageFile(t *testing.T) {
	testCases := []struct {
		expectError bool
		URL         string
	}{
		{expectError: true, URL: ""},
		{expectError: false, URL: "https://source.unsplash.com/random"},
		{expectError: false, URL: "https://thispersondoesnotexist.com/image"},
		{expectError: false, URL: "https://picsum.photos/400"},
		// {expectError: false, URL: "http://lorempixel.com/400/200"},
		{expectError: false, URL: "https://thiscatdoesnotexist.com/"},
	}
	for i, tc := range testCases {
		path := fmt.Sprintf("test-%d", i)
		err := ImageFile(path, tc.URL)
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