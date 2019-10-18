package service

import (
	"os"

	"github.com/ahmdrz/goinsta/v2"
	log "github.com/sirupsen/logrus"
)

type InstagramManager interface {
	FetchImageUrlsByTag(string) ([]string, error)
	Close() error
}

type InstagramClient struct {
	insta *goinsta.Instagram
}

func NewInstagramClient() (*InstagramClient, error) {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Info("error logging in")
		return nil, err
	}
	return &InstagramClient{
		insta: insta,
	}, nil
}

func (i *InstagramClient) FetchImageUrlsByTag(tag string) ([]string, error) {
	log.Info("fetching tag: ", tag)
	var results []string
	feedTag, err := i.insta.Feed.Tags(tag)
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

func (i *InstagramClient) Close() error {
	return i.insta.Logout()
}

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
