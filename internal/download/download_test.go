package download

import (
	"fmt"
	"testing"
)

func TestImageFile(t *testing.T) {
	testCases := []struct {
		URL string
	}{
		{URL: "https://source.unsplash.com/random"},
		{URL: "https://thispersondoesnotexist.com/image"},
		{URL: "https://picsum.photos/400"},
		// {URL: "http://lorempixel.com/400/200"},
		{URL: "https://thiscatdoesnotexist.com/"},
	}
	for i, testCase := range testCases {
		path := fmt.Sprintf("test-%d", i)
		err := ImageFile(path, testCase.URL)
		if err != nil {
			t.Error(fmt.Errorf("error on test case: %s: %s", testCase.URL, err))
		}
	}
}
