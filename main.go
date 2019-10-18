package main

import (
	"github.com/phantompunk/mosaic/handler"
	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

func main() {

	client, err := service.NewInstagramClient()
	if err != nil {
		log.Fatal("Failed to login")
	}
	defer client.Close()

	mosaic := handler.MosaicLambda{
		InstagramManager: client,
	}
	mosaic.LocalRequest("golang")
}
