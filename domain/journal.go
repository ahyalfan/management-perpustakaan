package domain

import (
	"context"
	"database/sql"
	"rest_api_sederhana/dto"
)

const (
	JournalStatusInProgress = "IN_PROGRESS"
	JournalStatusCompleted  = "COMPLETED"
)

type Journal struct {
	ID         string       `gorm:"primary_key"`
	BookID     string       `gorm:"column:book_id"`
	StockCode  string       `gorm:"column:stock_code"`
	CustomerId string       `gorm:"column:customer_id"`
	Status     string       `gorm:"column:status"`
	BorrowedAt sql.NullTime `gorm:"column:borrowed_at"`
	ReturnedAt sql.NullTime `gorm:"column:returned_at"`
	DueAt      sql.NullTime `gorm:"column:due_at"`
}

type JournalSearch struct {
	CustomerId string
	Status     string
}

type JournalRepository interface {
	Find(ctx context.Context, se JournalSearch) ([]Journal, error)
	FindById(ctx context.Context, id string) (Journal, error)
	Save(ctx context.Context, j *Journal) (string, error)
	Update(ctx context.Context, journal *Journal) error
}

type JournalService interface {
	Index(ctx context.Context, req JournalSearch) ([]dto.JournalData, error)
	Create(ctx context.Context, req dto.CreateJournalRequest) (string, error)
	Return(ctx context.Context, req dto.ReturnJournalRequest) error
}
