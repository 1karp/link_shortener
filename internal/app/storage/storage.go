package storage

import "sync"

type Storage interface {
	Set(shortURL, fullURL string)
	Get(shortURL string) string
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

func (s *URLStorage) Get(shortURL string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.urlMap[shortURL]
}
