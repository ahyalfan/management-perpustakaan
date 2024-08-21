package domain

import (
	"context"
	"rest_api_sederhana/dto"
	"time"
)

type Media struct {
	ID        string    `gorm:"primary_key;column:id"`
	Path      string    `gorm:"column:path"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type MediaRepository interface {
	Save(ctx context.Context, media *Media) (string, error)
	FindById(ctx context.Context, id string) (Media, error)
	FindByIds(ctx context.Context, id []string) ([]Media, error)
}

type MediaService interface {
	Create(ctx context.Context, req dto.CreatedMediaRequest) (dto.MediaData, error)
}
