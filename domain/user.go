package domain

import "context"

type User struct {
	ID       string `gorm:"primary_key"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}
