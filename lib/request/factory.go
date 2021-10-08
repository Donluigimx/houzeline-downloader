package request

import (
	"errors"
	"os"
)

type FactoryOpts struct {
	RequestType string
}

func NewFactoryOpts(requestType string) *FactoryOpts {
	if len(requestType) == 0 {
		requestType = os.Getenv("REQUEST_TYPE")
	}

	return &FactoryOpts{
		requestType,
	}
}

func NewRequest(factoryOpts *FactoryOpts) (InternalRequest, error) {
	if factoryOpts.RequestType == "http" {
		return &HttpRequest{}, nil
	}

	return nil, errors.New("No RequestType in options sent")
}
