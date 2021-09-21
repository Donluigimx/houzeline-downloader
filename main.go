package main

import (
	"downloader/lib/storage"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
  log.Info("Getting storage")
	fs, err := storage.CreateStorage(storage.FactoryOpts{})

	if err != nil {
		log.Fatal(err)
	}

	fs.Store("hola.txt", []byte("Hola diosito"))
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func mainFake() {
	resp, err := http.Get("https://www.casasyterrenos.com/propiedad/casa-renta-eulogio-parra-a-villa-senor-guadalajara-jal-2523534")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
