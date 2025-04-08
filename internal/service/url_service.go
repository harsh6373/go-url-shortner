package service

import (
	"errors"
	"time"

	"github.com/harsh6373/go-url-shortner/internal/model"
	"github.com/harsh6373/go-url-shortner/internal/repository"
	"github.com/harsh6373/go-url-shortner/internal/utils"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo}
}

func (s *URLService) Shorten(original, customSlug string, expiresAt *time.Time) (*model.URL, error) {
	slug := customSlug
	if slug == "" {
		slug = utils.GenerateSlug(6)
	}

	url := &model.URL{
		Slug:      slug,
		Original:  original,
		ExpiresAt: expiresAt,
	}

	if err := s.repo.CreateURL(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *URLService) Resolve(slug, userAgent string) (string, error) {
	url, err := s.repo.GetBySlug(slug)
	if err != nil {
		return "", err
	}

	if url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now()) {
		return "", errors.New("URL has expired")
	}

	s.repo.LogClick(&model.Click{
		Slug:      slug,
		UserAgent: userAgent,
	})

	return url.Original, nil
}

func (s *URLService) GetClickAnalytics(slug string) ([]model.Click, error) {
	return s.repo.GetClicksBySlug(slug)
}
