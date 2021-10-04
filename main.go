package main

import (
	"context"
	"crypto/sha256"
	"downloader/lib/storage"
	"fmt"
	"io/ioutil"
	"net/http"
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
	log.Info("Getting storage")
	fs, err := storage.CreateStorage(storage.NewFactoryOpts("", ""))

	if err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("PROPERTY_URL")
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	hashedUrl := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	fileName := fmt.Sprintf("%s.html", hashedUrl)
	log.WithField("file_name", fileName).Info("Storing file")
	fs.Store(fileName, body)
}

func mainFake() {
	resp, err := http.Get("https://www.casasyterrenos.com/propiedad/casa-renta-eulogio-parra-a-villa-senor-guadalajara-jal-2523534")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
