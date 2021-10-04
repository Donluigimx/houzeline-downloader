package storage

import (
	"errors"
	"os"
)

type FactoryOpts struct {
	StorageType string
	FullPath    string
}

func NewFactoryOpts(StorageType string, FullPath string) *FactoryOpts {
	if len(StorageType) == 0 {
		StorageType = os.Getenv("STORAGE_TYPE")
	}

	if len(FullPath) == 0 {
		FullPath = os.Getenv("STORAGE_PATH")
	}

	return &FactoryOpts{
		StorageType,
		FullPath,
	}
}

func CreateStorage(opts *FactoryOpts) (InternalStorage, error) {
	if opts.StorageType == "file" {
		if len(opts.FullPath) == 0 {
			return nil, errors.New("FullPath missing from options")
		}

		return FileStorage{opts.FullPath}, nil
	}

	return nil, errors.New("No StorageType sent on options")
}
