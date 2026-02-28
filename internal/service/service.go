package service

import (
	"url-encode/internal/storage"
)

type Service struct {
	store *storage.URLStore
}

func NewService(store *storage.URLStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateShortURL(originalURL string) string {
	if existing := s.store.FindByOriginal(originalURL); existing != "" {
		return existing
	}
	short := GenerateShortLink(originalURL)
	s.store.Save(short, originalURL)
	return short
}

func (s *Service) GetOriginalURL(shortURL string) (string, error) {
	return s.store.Get(shortURL)
}
