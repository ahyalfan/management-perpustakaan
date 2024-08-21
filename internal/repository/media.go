package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMedia(conn *gorm.DB) domain.MediaRepository {
	return &MediaRepository{db: conn}
}

// FindById implements domain.MediaRepository.
func (m *MediaRepository) FindById(ctx context.Context, id string) (media domain.Media, err error) {
	err = m.db.WithContext(ctx).Where("id =?", id).First(&media).Error
	return
}

// FindByIds implements domain.MediaRepository.
func (m *MediaRepository) FindByIds(ctx context.Context, id []string) (medias []domain.Media, err error) {
	err = m.db.WithContext(ctx).Where("id IN ?", id).Find(&medias).Error
	return
}

// Save implements domain.MediaRepository.
func (m *MediaRepository) Save(ctx context.Context, media *domain.Media) (string, error) {
	err := m.db.WithContext(ctx).Create(&media).Error
	if err != nil {
		return "", err
	}
	return media.ID, nil
}
