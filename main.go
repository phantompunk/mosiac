package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phantompunk/mosaic/handler"
	"github.com/phantompunk/mosaic/helper"
	"github.com/phantompunk/mosaic/service"
)

func main() {
	imageDownloader := service.NewInstagramDownloader(helper.InstagramHelper)

	mosaicLambda := handler.MosaicLambda{
		Name:         "rigo",
		ImageManager: imageDownloader,
	}

	lambda.Start(mosaicLambda.HandleRequest)
}

// --------------------------------------------------------
// import (
// 	"bytes"
// 	"errors"
// 	"io"
// 	"net/http"
// 	"os"

// 	log "github.com/sirupsen/logrus"
// 	"gopkg.in/ahmdrz/goinsta.v2"
// )

// func main() {
// 	// 1. Query for Insta images
// 	insta := goinsta.New(
// 		"phantompunk_",
// 		"rigo?Insta13",
// 	)
// 	if err := insta.Login(); err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	defer insta.Logout()

// 	feedTag, err := insta.Feed.Tags("golang")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	// 2. Collect image urls
// 	itemUrls := make([]string, 1)
// 	for i := 0; i <= 10; i++ {
// 		itemURL := feedTag.Images[i].Images.GetBest()
// 		if itemURL == "" {
// 			log.Infof("Item URL is empty: %s, index: %d", itemURL, i)
// 		} else {
// 			log.Infof("Item URL added: %s, index: %d", itemURL, i)
// 			itemUrls = append(itemUrls, itemURL)
// 		}
// 	}
// 	// for _, url := range itemUrls {
// 	// 	println(url)
// 	// }

// 	// 3. Download images
// 	// for index, url := range itemUrls {
// 	// 	if url != "" {
// 	// 		log.Infof("index: %d, value %s", index, url)
// 	// 	}
// 	// }

// 	// log.Infof("Print 1st url: \v\n", itemUrls[0:])
// 	// log.Infof("Print 1st url: \v\n", itemUrls[3])
// 	// log.Infof("Print 1st url: \v\n", itemUrls[6])
// 	// log.Infof("Print 1st url: %s", itemUrls[13])

// 	_, err = downloadImage(itemUrls[9])
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	// 4. Gather images in list
// 	// 5. Combine images into one
// 	// 6. return image
// }

// func downloadImage(url string) ([]byte, error) {

// 	println("download: " + url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Error("error fetching image")
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		return nil, errors.New(response.Status)
// 	}

// 	file, err := os.Create("test.jpg")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var data bytes.Buffer
// 	_, err = io.Copy(file, response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data.Bytes(), nil
// }
