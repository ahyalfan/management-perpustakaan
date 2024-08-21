package domain

import (
	"context"
	"rest_api_sederhana/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
