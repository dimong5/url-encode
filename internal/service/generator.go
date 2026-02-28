package service

import (
	"crypto/md5"
	"encoding/hex"
)

const shortLinkLength = 10

func GenerateShortLink(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))
	hashStr := hex.EncodeToString(hash[0:16])
	return hashStr[:shortLinkLength]
}