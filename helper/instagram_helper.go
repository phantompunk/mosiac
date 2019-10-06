package helper

import (
	"os"

	"github.com/ahmdrz/goinsta/v2"
)

// InstagramHelper handles all the instagram dependencies and image fetching
var InstagramHelper = func(queryTag string) ([]string, error) {

	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		return nil, err
	}
	defer insta.Logout()

	feedTag, err := insta.Feed.Tags(queryTag)
	if err != nil {
		return nil, err
	}

	foundUrls := make([]string, 20)
	for _, item := range feedTag.Images {
		url := item.Images.GetBest()
		foundUrls = append(foundUrls, url)
	}

	return foundUrls, nil
}
