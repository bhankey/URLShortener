package store

import (
	"database/sql"
	"fmt"
)

// urlRepository current realisation of repository
type urlRepository struct {
	s *store
}

// URLRepository defines the interface for repository to implement
type URLRepository interface {
	Create(originalURL string) (shortUrl string, err error)
	Get(shootUrl string) (originalUrl string, err error)
}

// Create the url in store and generates short code if it does not exist already. returns short code for url
func (r *urlRepository) Create(u string) (string, error) {
	shortURL, err := r.getShortURL(u)
	if err == nil {
		return shortURL, nil
	} else if err != ErrNoRecords {
		return "", err
	}
	shortURL = RandomStringGenerator(LengthURL)
	isCollision, err := r.isShortURLExist(shortURL)
	if err != nil {
		return "", err
	}
	if isCollision {
		for i := 0; isCollision && i < 20; i++ {
			shortURL = RandomStringGenerator(LengthURL)
			isCollision, err = r.isShortURLExist(shortURL)
			if err != nil {
				return "", err
			}
		}
		if isCollision {
			return "", fmt.Errorf("could not create short url, because of collisions, try again later")
		}
	}
	var originalURL string
	if err := r.s.db.QueryRow("INSERT INTO url VALUES ($1, $2) RETURNING original_url", shortURL, u).Scan(&originalURL); err != nil {
		return "", err
	}
	return shortURL, nil
}

// Get the original URL by it's short code
func (r *urlRepository) Get(shortUrl string) (string, error) {
	var url string
	if err := r.s.db.QueryRow(
		"SELECT original_url FROM url WHERE short_url = $1",
		shortUrl).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNoRecords
		}
		return "", err
	}
	return url, nil
}

// isShortURLExist checks the existing of short code in the store
func (r *urlRepository) isShortURLExist(shortURL string) (bool, error) {
	var isShortURL bool
	if err := r.s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM url WHERE short_url = $1)",
		shortURL).Scan(&isShortURL); err != nil {
		return false, err
	}
	return isShortURL, nil
}

// getShortURL returns short code if original URL exists in the store
func (r *urlRepository) getShortURL(OriginalURL string) (string, error) {
	var shortURL string
	if err := r.s.db.QueryRow("SELECT short_url FROM url WHERE original_url = $1",
		OriginalURL).Scan(&shortURL); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNoRecords
		}
		return "", err
	}
	return shortURL, nil
}
