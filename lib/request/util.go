package request

import (
	"crypto/sha256"
	"fmt"
)

func HashUrl(url string) string {
	hashedUrl := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))

	return hashedUrl
}
