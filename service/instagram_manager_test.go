package service

import (
	"reflect"
	"testing"

	"github.com/ahmdrz/goinsta/v2"
)

func TestFetchUrls(t *testing.T) {
	client := MockInstagramClient{}
	results, _ := client.FetchUrlsByTag("fake")
	want := []string{
		"https://insta.fbcdn.net/v/t51/e35/p1080x1080/7630_n.jpg",
		"https://insta.fbcdn.net/v/t56/e25/p1080x1080/47530_n.jpg",
	}

	if results == nil {
		t.Errorf("Fetched Urls is nil")
	}

	if !reflect.DeepEqual(results, want) {
		t.Errorf("Results not equal, got %v want %v", results, want)
	}
}

type MockInstagramClient struct {
	client *goinsta.Instagram
}

func (m *MockInstagramClient) FetchUrlsByTag(tag string) ([]string, error) {
	return []string{
		"https://insta.fbcdn.net/v/t51/e35/p1080x1080/7630_n.jpg",
		"https://insta.fbcdn.net/v/t56/e25/p1080x1080/47530_n.jpg",
	}, nil
}
