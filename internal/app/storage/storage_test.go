package storage

import (
	"testing"
)

func TestURLStorage_SetAndGet(t *testing.T) {
	storage := NewStorage()

	shortURL := "abc"
	fullURL := "https://example.com"
	storage.Set(shortURL, fullURL)

	result := storage.Get(shortURL)
	if result != fullURL {
		t.Errorf("Get(%s) = %s; want %s", shortURL, result, fullURL)
	}
}

func TestURLStorage_GetNonExistent(t *testing.T) {
	storage := NewStorage()

	shortURL := "nonexistent"
	result := storage.Get(shortURL)
	if result != "" {
		t.Errorf("Get(%s) = %s; want empty string", shortURL, result)
	}
}
