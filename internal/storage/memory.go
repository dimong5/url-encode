package storage

import "errors"

var ErrNotFound = errors.New("not found")

type URLStore struct {
	urls map[string]string // [short]original
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

func (s *URLStore) Save(short, original string) {
	s.urls[short] = original
}

func (s *URLStore) Get(short string) (string, error) {
	url, ok := s.urls[short]
	if !ok {
		return "", ErrNotFound
	}
	return url, nil
}

func (s *URLStore) FindByOriginal(original string) string {
	for short, orig := range s.urls {
		if orig == original {
			return short
		}
	}
	return ""
}
