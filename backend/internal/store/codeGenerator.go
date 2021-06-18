package store

import "math/rand"

// LengthURL - length of generated code
const (
	LengthURL = 10
)

const alphabetForGeneration = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// RandomStringGenerator generates short code for URL
func RandomStringGenerator(length int) string {
	res := make([]byte, length)
	for i := range res {
		res[i] = alphabetForGeneration[rand.Int63()%int64(len(alphabetForGeneration))]
	}
	return string(res)
}
