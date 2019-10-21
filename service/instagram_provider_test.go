package service

import (
	"reflect"
	"testing"

	"github.com/ahmdrz/goinsta/v2"
)

func TestFetchUrls(t *testing.T) {
	mockClient := MockInstagramClient{}
	mockFeed := MockInstagramFeed{}
	provider := &InstagramProvider{
		Client: mockClient,
		Feed:   mockFeed,
	}

	results, _ := provider.SearchByTag("fake")
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

func TestDownloadImage(t *testing.T) {
	results := []string{
		"https://insta.fbcdn.net/v/t51/e35/p1080x1080/7630_n.jpg",
		"https://insta.fbcdn.net/v/t56/e25/p1080x1080/47530_n.jpg",
	}

	provider := &InstagramProvider{}

	for _, url := range results {
		img, err := provider.DownloadImage(url)
		if err != nil {
			t.Fatalf("Image download failed")
		}
		if img != nil {
		}
	}
}

type MockInstagramClient struct{}

type MockInstagramFeed struct{}

func (m MockInstagramClient) Login() error {
	return nil
}

func (m MockInstagramClient) Logout() error {
	return nil
}

func (m MockInstagramFeed) Tags(tag string) (*goinsta.FeedTag, error) {
	return &goinsta.FeedTag{
		Images: []goinsta.Item{
			goinsta.Item{
				Images: goinsta.Images{
					Versions: []goinsta.Candidate{
						goinsta.Candidate{
							Width:  100,
							Height: 100,
							URL:    "https://insta.fbcdn.net/v/t51/e35/p1080x1080/7630_n.jpg",
						},
					},
				},
			},
			goinsta.Item{
				Images: goinsta.Images{
					Versions: []goinsta.Candidate{
						goinsta.Candidate{
							Width:  100,
							Height: 100,
							URL:    "https://insta.fbcdn.net/v/t56/e25/p1080x1080/47530_n.jpg",
						},
					},
				},
			},
		},
	}, nil
}
