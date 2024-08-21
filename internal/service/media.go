package service

import (
	"context"
	"path"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/config"

	"github.com/google/uuid"
)

type MediaService struct {
	conf            *config.Config
	mediaRepository domain.MediaRepository
}

func NewMedia(conf *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &MediaService{mediaRepository: mediaRepository, conf: conf}
}

// Create implements domain.MediaService.
func (m *MediaService) Create(ctx context.Context, req dto.CreatedMediaRequest) (dto.MediaData, error) {
	id, err := m.mediaRepository.Save(ctx, &domain.Media{
		ID:   uuid.NewString(),
		Path: req.Path,
	})
	if err != nil {
		return dto.MediaData{}, err
	}
	url := path.Join(m.conf.Server.Asset, req.Path)
	return dto.MediaData{ID: id, Path: req.Path, Url: url}, nil
}
