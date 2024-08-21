package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUser(conn *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	err = ur.db.WithContext(ctx).Take(&user, "email = ?", email).Error
	return
}
