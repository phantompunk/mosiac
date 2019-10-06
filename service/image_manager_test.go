package service

import (
	"reflect"
	"testing"
)

func Test_instagramDownloader_FetchImageUrlsByHashtag(t *testing.T) {

	tests := []struct {
		name        string
		fetcherFunc InstagramFetcher
		hashtag     string
		want        []string
		wantErr     bool
	}{
		{
			"no tagged images founds",
			mockHelperFoundNothing,
			"golang",
			nil,
			false,
		},
		{
			"tagged images found",
			mockHelperFoundImages,
			"golang",
			[]string{"http://google.com/j.jpg", "http://rigomoran.com/t.jpg"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insta := instagramDownloader{
				fetcherFunc: tt.fetcherFunc,
			}
			got, err := insta.FetchImageUrlsByHashtag(tt.hashtag)
			if (err != nil) != tt.wantErr {
				t.Errorf("instagramDownloader.FetchImageUrlsByHashtag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instagramDownloader.FetchImageUrlsByHashtag() = %v, want %v", got, tt.want)
			}
		})
	}
}

var mockHelperFoundNothing = func(queryTag string) ([]string, error) {
	var empty []string
	return empty, nil
}

var mockHelperFoundImages = func(queryTag string) ([]string, error) {
	return []string{"http://google.com/j.jpg", "http://rigomoran.com/t.jpg"}, nil
}
