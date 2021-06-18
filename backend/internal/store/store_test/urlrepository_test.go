package store_test

import (
	"URLShortener/internal/store"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

// testDBConfig - config for connecting to testing DB
var testDBConfig = store.Config{
	Host:     "localhost",
	Port:     5432,
	User:     "test",
	Password: "test",
	DBName:   "URLShortenerTest",
}

func TestURLRepository_CreateGet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s, teardown := store.TestStore(t, &testDBConfig)
	defer teardown("url")

	originalURL := "https://yandex.ru"
	u, err := s.URL().Create(originalURL)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	u1, err := s.URL().Create(originalURL)
	assert.Equal(t, u, u1)
	assert.NoError(t, err)

	o, err := s.URL().Get(u)
	assert.NoError(t, err)
	assert.Equal(t, o, originalURL)

	o, err = s.URL().Get("somecode")
	assert.Error(t, err)
	assert.Empty(t, o)
}

func TestURLRepository_CreateGetRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	urls := make([]string, 10000)
	for i := range urls {
		urls[i] = "https://example.com/" + store.RandomStringGenerator(10)
	}
	s, teardown := store.TestStore(t, &testDBConfig)
	defer teardown("url")
	shortUrls := make([]string, 10000)
	for i, url := range urls {
		shortUrl, err := s.URL().Create(url)
		assert.NoError(t, err)
		shortUrls[i] = shortUrl
	}

	for i, shortUrl := range shortUrls {
		url, err := s.URL().Get(shortUrl)
		assert.NoError(t, err)
		assert.Equal(t, urls[i], url)
	}
}
