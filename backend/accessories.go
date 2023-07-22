package backend

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

func makeShortcode(longUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(longUrl))
	hashedUrl := hex.EncodeToString(hasher.Sum(nil))

	return hashedUrl
}

func isURL(ur string) bool {
	u, err := url.Parse(ur)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
