package repository

import (
	"github.com/harsh6373/go-url-shortner/internal/model"
	"gorm.io/gorm"
)

type URLRepository interface {
	CreateURL(url *model.URL) error
	GetBySlug(slug string) (*model.URL, error)
	LogClick(click *model.Click) error
}

type urlRepo struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepo{db}
}

func (r *urlRepo) CreateURL(url *model.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepo) GetBySlug(slug string) (*model.URL, error) {
	var url model.URL
	err := r.db.Where("slug = ?", slug).First(&url).Error
	return &url, err
}

func (r *urlRepo) LogClick(click *model.Click) error {
	return r.db.Create(click).Error
}
