package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phantompunk/mosaic/handler"
	"github.com/phantompunk/mosaic/service"
	log "github.com/sirupsen/logrus"
)

func main() {

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}

	provider, err := service.NewInstagramProvider()
	if err != nil {
		log.Fatal("Failed to login")
	}
	defer provider.Client.Logout()

	transformer := service.NewCanvas(3, 300)

	mosaic := handler.MosaicLambda{
		ImageProvider: provider,
		Transformer:   transformer,
	}

	if env == "dev" {
		lambda.Start(mosaic.HandleRequest)
	} else {
		mosaic.LocalRequest("worldjellyfishday")
	}
}
