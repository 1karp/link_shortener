package storage

var storage = make(map[string]string, 100)

func Set(shortURL, fullURL string) {
	storage[shortURL] = fullURL
}

func Get(shortURL string) string {
	return storage[shortURL]
}
