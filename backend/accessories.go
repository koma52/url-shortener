package backend

import (
	"crypto/md5"
	"encoding/hex"
)

func makeShortcode(longUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(longUrl))
	hashedUrl := hex.EncodeToString(hasher.Sum(nil))

	return hashedUrl
}
