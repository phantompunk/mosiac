package handler

import (
	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	Name             string
	InstagramManager service.InstagramManager
}

// HandleRequest is the main entry point for the lambda function
func (m *MosaicLambda) HandleRequest(searchTag string) (int, error) {
	// 1. Fetch photo urls
	log.Info("Start image fetch for ", searchTag)
	return 0, nil
}

// LocalRequest is the main entry point for the local execution
func (m *MosaicLambda) LocalRequest(searchTag string) (string, error) {
	log.Info("starting local run")

	// 1. Fetch photo urls
	results, _ := m.InstagramManager.FetchImageUrlsByTag(searchTag)
	log.Info(len(results), " images found")
	// 2. Transform photos
	// 2a. Create mosaic rectangle
	// 2b. Download photos
	// 2c. Place phot

	return "", nil
}
