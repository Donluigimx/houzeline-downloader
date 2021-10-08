package request

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpRequest struct{}

func (h *HttpRequest) Get(url string) (GetResponse, error) {
	log.WithFields(log.Fields{
		"url": url,
	}).Info("Requesting via HTTP/HTTPS")

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	key := HashUrl(url)

	log.WithFields(log.Fields{
		"key": key,
	}).Info("Returning Body")

	return GetResponse{body, key}, nil
}
