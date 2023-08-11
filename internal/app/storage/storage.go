package storage

import (
	"errors"
	"sync"
)

type Storage interface {
	Set(shortURL, fullURL string)
	Get(shortURL string) (string, error)
}

type URLStorage struct {
	urlMap map[string]string
	mutex  sync.Mutex
}

func NewStorage() *URLStorage {
	return &URLStorage{
		urlMap: make(map[string]string),
	}
}

func (s *URLStorage) Set(shortURL, fullURL string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.urlMap[shortURL] = fullURL
}

func (s *URLStorage) Get(shortURL string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if url, exists := s.urlMap[shortURL]; exists {
		return url, nil
	}

	return "", errors.New("short URL not found")
}
