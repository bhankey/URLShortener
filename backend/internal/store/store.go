package store

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

// store current realisation of storing data
type store struct {
	config        *Config
	db            *sql.DB
	URLRepository URLRepository
}

// Store defines the interface for storages to implement
type Store interface {
	Open() error
	Close()
	URL() URLRepository
}

var (
	ErrNoRecords       = errors.New("no record")
	ErrURLAlreadyExist = errors.New("URL was already shorted")
)

// New returns interface with current implementation of store
func New(config *Config) (Store, error) {
	s := &store{
		config: config,
	}
	err := s.Open()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Open connects to database and creates table with URLs if it doesn't exist
func (s *store) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.config.Host, s.config.Port, s.config.User, s.config.Password, s.config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	if _, err := s.db.Exec("CREATE TABLE IF NOT EXISTS url ( short_url text NOT NULL PRIMARY KEY, original_url text NOT NULL UNIQUE);"); err != nil {
		return err
	}
	return nil
}

// Close connection to database
func (s *store) Close() {
	_ = s.db.Close()
}

// URL returns URLRepository
func (s *store) URL() URLRepository {
	if s.URLRepository != nil {
		return s.URLRepository
	}
	s.URLRepository = &urlRepository{
		s: s,
	}
	return s.URLRepository
}
