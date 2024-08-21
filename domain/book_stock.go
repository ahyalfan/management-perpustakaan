package domain

import (
	"context"
	"database/sql"
	"rest_api_sederhana/dto"
)

const (
	BookStockStatusBorrowed  = "BORROWED"
	BookStockStatusAvailable = "AVAILABLE"
)

type BookStock struct {
	ID         int            `gorm:"primary_key;column:id;auto_increment"`
	BookId     string         `gorm:"column:book_id"`
	Code       string         `gorm:"column:code"`
	Status     string         `gorm:"column:status"`
	BorrowerId sql.NullString `gorm:"borrower_id"`
	BorrowedAt sql.NullTime   `gorm:"borrowed_at"`
}

type BookStockRepository interface {
	FindByBookId(ctx context.Context, id string) ([]BookStock, error)
	FindByCode(ctx context.Context, code string) (BookStock, error)
	Save(ctx context.Context, bookStock []BookStock) error

	Update(ctx context.Context, stock *BookStock) error
	DeleteByBookId(ctx context.Context, id int) error
	DeleteByCodes(ctx context.Context, code []string) error
}

type BookStockService interface {
	Create(ctx context.Context, req dto.CreateBookStockRequest) error
	Delete(ctx context.Context, req dto.DeleteBookStockRequest) error
}
