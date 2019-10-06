package handler

import "github.com/phantompunk/mosaic/service"

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	Name         string
	ImageManager service.ImageManager
}

// HandleRequest is the main entry point for the lambda function
func (m *MosaicLambda) HandleRequest() (int, error) {
	// 1. Fetch photo urls
	m.ImageManager.FetchImageUrlsByHashtag("")
	// 2. Transform photos
	// 2a. Create mosaic rectangle
	// 2b. Download photos
	// 2c. Place photo
	return 0, nil
}
