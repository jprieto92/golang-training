package websiteChecker

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{"https://google.com",
		"https://blog.gypsydave5.com",
		"https://furhurterwe.geds",
	}

	want := map[string]bool{"https://google.com": true,
		"https://blog.gypsydave5.com": true,
		"https://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
