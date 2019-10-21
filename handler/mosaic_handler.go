package handler

import (
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"os"

	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	ImageProvider *service.InstagramProvider
	Transformer   Transformer
}

type MosaicRequest struct {
	Hashtag string `json:"hashtag,omitempty"`
}

type MosaicResponse struct {
	Key string
}

type Transformer interface {
	Merge([]image.Image) image.RGBA
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

	// 2. Create a 2x2 Rectangle
	var width, height = 320, 320
	rect := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: 5 * width,
			Y: 5 * height,
		},
	})

	jobs := make(chan string, 5)
	images := make(chan image.Image, 5)

	for w := 1; w <= 25; w++ {
		go worker(w, jobs, images, rect)
	}

	for j := 1; j <= 25; j++ {
		jobs <- results[j]
	}
	close(jobs)

	for a := 1; a <= 25; a++ {
		<-images
	}

	out, err := os.Create("./merged.jpg")
	if err != nil {
	}

	var opt jpeg.Options
	opt.Quality = 80

	jpeg.Encode(out, rect, &opt)
	// img, err := service.DownloadImage(results[0])
	// if err != nil {
	// }

	// out, err := os.Create("./output.png")
	// if err != nil {
	// }

	// err = png.Encode(out, img)
	// 2. Transform photos
	// 2a. Create mosaic rectangle
	// 2b. Download photos
	// 2c. Place phot

	return searchTag, nil
}

func worker(id int, jobs <-chan string, images chan<- image.Image, rect *image.RGBA) {
	for j := range jobs {
		img, _ := service.DownloadImage(j, id)
		merge(id, img, rect)
		// draw.Draw(rect, img.Bounds(), img, image.Point{}, draw.Src)
		images <- img
	}
}

func merge(id int, img image.Image, rect *image.RGBA) {
	x := id % 5
	y := math.Round(float64(id / 5))
	minPoint := image.Point{x * 320, int(y) * 320}
	maxPoint := minPoint.Add(image.Point{320, 320})
	nextRect := image.Rectangle{minPoint, maxPoint}

	draw.Draw(rect, nextRect, img, image.Point{}, draw.Src)
}
