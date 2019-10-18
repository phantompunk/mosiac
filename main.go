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
		env = "dev"
	}
	client, err := service.NewInstagramClient()
	if err != nil {
		log.Fatal("Failed to login")
	}
	defer client.Close()

	mosaic := handler.MosaicLambda{
		InstagramManager: client,
	}

	if env == "dev" {
		lambda.Start(mosaic.HandleRequest)
	} else {
		mosaic.LocalRequest("golang")
	}
}
