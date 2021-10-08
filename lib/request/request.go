package request

type GetResponse struct {
	Body []byte
	Key  string
}

type InternalRequest interface {
	Get(url string) (GetResponse, error)
}
