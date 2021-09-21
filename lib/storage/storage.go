package storage

type InternalStorage interface {
	Store(key string, data []byte)
}
