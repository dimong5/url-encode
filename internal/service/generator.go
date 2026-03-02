package service

import (
	"crypto/md5"
)

//const shortLinkLength = 10

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"


func GenerateShortLink(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))
	//hashStr := hex.EncodeToString(hash[:])
	//return hashStr[:shortLinkLength]
	result := make([]byte, 10)
		for i := 0; i < 10; i++ {
		result[i] = charset[int(hash[i])%63]
	}
	return string(result)
}