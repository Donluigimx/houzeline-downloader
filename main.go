package main

import (
	"context"
	"downloader/lib/request"
	"downloader/lib/storage"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) {
	// This would not work
	lambda.Start(handler)
}

func main() {
	storageFactoryOpts := storage.NewFactoryOpts("", "")
	log.WithFields(log.Fields{
		"factory_opts": storageFactoryOpts,
	}).Info("Storage Factory Opts created")

	log.Info("Getting storage")
	fs, err := storage.NewStorage(storageFactoryOpts)

	if err != nil {
		log.Fatal(err)
	}

	requestFactoryOpst := request.NewFactoryOpts("")
	log.WithFields(log.Fields{
		"factory_opts": requestFactoryOpst,
	}).Info("Request Factory Opts created")

	log.Info("Getting request")
	r, err := request.NewRequest(requestFactoryOpst)

	if err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("PROPERTY_URL")
	response, err := r.Get(url)

	log.WithFields(log.Fields{
		"key": response.Key,
	}).Info("Storing file")

	fs.Store(response.Key, response.Body)
}
