package handler

import (
	"image"
	"math"

	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	ImageProvider *service.InstagramProvider
	Transformer   *service.Transformer
}

type MosaicRequest struct {
	Hashtag string `json:"hashtag,omitempty"`
}

type MosaicResponse struct {
	Key string
}

// HandleRequest is the main entry point for the lambda function
func (m *MosaicLambda) HandleRequest(req MosaicRequest) (MosaicResponse, error) {
	// 1. Fetch photo urls
	log.Info("Start image fetch for ", req.Hashtag)
	return MosaicResponse{
		Key: req.Hashtag,
	}, nil
}

// LocalRequest is the main entry point for the local execution
func (m *MosaicLambda) LocalRequest(searchTag string) (string, error) {
	log.Info("starting local run")

	// 1. Fetch photo urls
	results, _ := m.ImageProvider.SearchByTag(searchTag)
	log.Info(len(results), " images found")

	// 2. Create a 2x2 RectangleS
	size := m.Transformer.Size
	grids := int(math.Pow(float64(size), 2))
	log.Info("Size before:", size)

	// 3. Download images and merge
	jobs := make(chan string, size)
	images := make(chan image.Image, size)

	for w := 1; w <= grids; w++ {
		go worker(*m, w, jobs, images, m.Transformer.Canvas)
	}
	for j := 1; j <= grids; j++ {
		jobs <- results[j]
	}
	close(jobs)
	for a := 1; a <= grids; a++ {
		<-images
	}

	m.Transformer.Export()

	return searchTag, nil
}

func worker(m MosaicLambda, id int, jobs <-chan string, images chan<- image.Image, rect *image.RGBA) {
	for j := range jobs {
		img, err := service.DownloadImage(j, id)
		if err != nil {
			log.Error("Image couldn't download compeletly")
		}
		// service.transformer.combine()
		// merge(id-1, img, rect)
		m.Transformer.Merge(img)
		// draw.Draw(rect, img.Bounds(), img, image.Point{}, draw.Src)
		images <- img
	}
}
