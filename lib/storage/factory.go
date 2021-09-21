package storage

import (
	"errors"
	"os"
)

type FactoryOpts struct {
	StorageType string
	FullPath    string
}

func CreateStorage(opts FactoryOpts) (InternalStorage, error) {
	if len(opts.StorageType) == 0 {
		opts.StorageType = os.Getenv("STORAGE_TYPE")
	}

	if opts.StorageType == "file" {
		if len(opts.FullPath) == 0 {
			opts.FullPath = os.Getenv("STORAGE_PATH")

			if len(opts.FullPath) == 0 {
				return nil, errors.New("FullPath missing from options")
			}
		}

		return FileStorage{opts.FullPath}, nil
	}

	return nil, errors.New("No StorageType sent on options")
}
