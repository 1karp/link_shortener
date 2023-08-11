package storage

import (
	"testing"
)

func TestURLStorage_SetAndGet(t *testing.T) {
	storage := NewStorage()

	shortURL := "abc"
	fullURL := "https://example.com"
	storage.Set(shortURL, fullURL)

	result, err := storage.Get(shortURL)
	if err != nil {
		t.Errorf("Error getting URL: %v", err)
	}

	if result != fullURL {
		t.Errorf("Get(%s) = %s; want %s", shortURL, result, fullURL)
	}
}

func TestURLStorage_GetNonExistent(t *testing.T) {
	storage := NewStorage()

	shortURL := "nonexistent"
	result, err := storage.Get(shortURL)
	if err == nil {
		t.Errorf("Expected error getting nonexistent URL, but got none")
	}

	if result != "" {
		t.Errorf("Get(%s) = %s; want empty string", shortURL, result)
	}
}
