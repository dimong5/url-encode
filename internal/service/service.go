package service

import (
	//"url-encode/internal/storage"
)

type Service struct {
		store interface {
			Save(short, original string) error
			Get(short string) (string, error)
			FindByOriginal(original string) string
	}
}

func NewService(		store interface {
			Save(short, original string) error
			Get(short string) (string, error)
			FindByOriginal(original string) string
	}) *Service {
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
