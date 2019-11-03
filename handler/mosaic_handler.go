package handler

import (
	"image"
	"image/draw"
	"math"

	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	ImageProvider *service.InstagramProvider
	Transformer   service.Transformer
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

	// 2. Create a 2x2 Rectangle
	// var width = 320
	// rectangle := image.NewRGBA(image.Rectangle{
	// 	Max: image.Point{
	// 		X: 5 * width,
	// 		Y: 5 * height,
	// 	},
	// })

	// m.Transformer = &Transformer{}
	log.Info("Size before:", m.Transformer.Size)

	// 3. Download images and merge
	jobs := make(chan string, 5)
	images := make(chan image.Image, 5)

	for w := 1; w <= 25; w++ {
		go worker(*m, w, jobs, images, m.Transformer.Canvas)
	}
	for j := 1; j <= 25; j++ {
		jobs <- results[j]
	}
	close(jobs)
	for a := 1; a <= 25; a++ {
		<-images
	}

	// 4. Return the combined image
	// out, err := os.Create("./merged.jpg")
	// if err != nil {
	// }
	// var opt jpeg.Options
	// opt.Quality = 80
	// jpeg.Encode(out, rectangle, &opt)
	// service.transformer.Export()
	m.Transformer.Export()

	return searchTag, nil
}

func worker(m MosaicLambda, id int, jobs <-chan string, images chan<- image.Image, rect *image.RGBA) {
	for j := range jobs {
		img, _ := service.DownloadImage(j, id)
		// service.transformer.combine()
		// merge(id-1, img, rect)
		log.Info("Size:", m.Transformer.Size)
		m.Transformer.Merge(img)
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
