package storage

import (
	"io/ioutil"
	"path"
)

type FileStorage struct {
	fullPath string
}

func (f FileStorage) Store(key string, data []byte) {
	joinedPath := path.Join(f.fullPath, key)

	ioutil.WriteFile(joinedPath, data, 0777)

	return
}
