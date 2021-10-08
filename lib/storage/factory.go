package storage

import (
	"errors"
	"os"
)

type FactoryOpts struct {
	StorageType string
	FullPath    string
}

func NewFactoryOpts(storageType string, fullPath string) *FactoryOpts {
	if len(storageType) == 0 {
		storageType = os.Getenv("STORAGE_TYPE")
	}

	if len(fullPath) == 0 {
		fullPath = os.Getenv("STORAGE_PATH")
	}

	return &FactoryOpts{
		storageType,
		fullPath,
	}
}

func NewStorage(opts *FactoryOpts) (InternalStorage, error) {
	if opts.StorageType == "file" {
		if len(opts.FullPath) == 0 {
			return nil, errors.New("FullPath missing from options")
		}

		return FileStorage{opts.FullPath}, nil
	}

	return nil, errors.New("No StorageType sent on options")
}
