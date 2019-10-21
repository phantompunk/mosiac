package service

import (
	"os"

	"github.com/ahmdrz/goinsta/v2"
	log "github.com/sirupsen/logrus"
)

// InstagramClient represents a goinsta client
type InstagramClient interface {
	Login() error
	Logout() error
}

// InstagramFeed represents a goinsta feed struct
type InstagramFeed interface {
	Tags(tag string) (*goinsta.FeedTag, error)
}

// InstagramProvider represents a mockable goinsta service
type InstagramProvider struct {
	Client InstagramClient
	Feed   InstagramFeed
}

// SearchByTag represents a tag search
func (i *InstagramProvider) SearchByTag(tag string) ([]string, error) {
	log.Info("fetching tag: ", tag)

	var results []string
	feedTag, err := i.Feed.Tags(tag)
	if err != nil {
		return nil, err
	}

	for _, item := range feedTag.Images {
		if item.Images.Versions != nil {
		}
		url := GetSmallestImage(item.Images)
		if url != "" {
			results = append(results, url)
		}
	}
	return results, nil
}

// NewInstagramProvider creates a new Instagram structure
func NewInstagramProvider() (*InstagramProvider, error) {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Info("error logging in")
		return nil, err
	}
	return &InstagramProvider{
		Client: insta,
		Feed:   insta.Feed,
	}, nil
}

// GetSmallestImage returns the smallest sized image available
func GetSmallestImage(img goinsta.Images) string {
	temp := ""
	for _, v := range img.Versions {
		temp = v.URL
		th, tw := v.Height, v.Width
		if v.Width < tw || v.Height < th {
			temp = v.URL
			th, tw = v.Height, v.Width
		}
	}
	smallest := temp
	return smallest
}

// Example Code
////////////////////////////////////////////
// func main() {

// 	client, err := service.NewInstagramClient()
// 	if err != nil {
// 		log.Fatal("Failed to login")
// 	}
// 	defer client.Close()

// 	list, _ := client.FetchUrlsByTag("golang")
// 	fmt.Println(list[:5])
// }
// type ImageManager interface {
// 	FetchImageUrlsByTag() ([]string, error)
// }
