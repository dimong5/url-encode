package storage

import "sync"

type URLStore struct {
	urls sync.Map
}

func NewURLStore() *URLStore {
	return &URLStore{}
}

func (s *URLStore) Save(short, original string) error {
	s.urls.Store(short, original)
	return nil
}

func (s *URLStore) Get(short string) (string, error) {
	url, ok := s.urls.Load(short)
	if !ok {
		return "", ErrNotFound
	}
	return url.(string), nil
}

func (s *URLStore) FindByOriginal(original string) string {
	var result string
	
	callback := func(key, value any) bool {
		if value == original {
			result = key.(string)
			return false
		}
		return true
	}
	
	s.urls.Range(callback)
	return result
}
