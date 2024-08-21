package service

import (
	"context"
	"errors"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, ur domain.UserRepository) domain.AuthService {
	return &AuthService{
		conf:           cnf,
		userRepository: ur,
	}
}

// Login implements domain.AuthService.
func (as *AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := as.userRepository.FindByEmail(ctx, req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}
	if err != nil {
		return dto.AuthResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// kita coba bandingkan password yg sudah di hash sama password yg dimasukan user
	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	// bikin sebuah token
	claims := jwt.MapClaims{
		// kita bebes mau masukin apa, asal jangan sensitif
		"id": user.ID,
		// waktu expire nya
		"exp": time.Now().Add(time.Duration(as.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	// kita mau pilih signaturenya yg mana, biasanya sih hs256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// kemudian kita pastikan tokennya terbuat dengan key yg sudah ada di env
	tokenString, err := token.SignedString([]byte(as.conf.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, errors.New("authentication failed")
	}

	return dto.AuthResponse{
		Token: tokenString,
	}, nil
}
