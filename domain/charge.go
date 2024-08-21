package domain

import (
	"context"
	"time"
)

type Charge struct {
	ID           string    `gorm:"primary_key"`
	UserID       string    `gorm:"column:user_id"`
	JournalID    string    `gorm:"column:journal_id"`
	DaysLate     int       `gorm:"column:days_late"`
	DailyLateFee int       `gorm:"column:daily_late_fee"`
	Total        int       `gorm:"column:total"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type ChargeRepository interface {
	Save(ctx context.Context, r *Charge) (string, error)
	FindByUserID(ctx context.Context, userID string) ([]Charge, error)
	FindByJournalID(ctx context.Context, journalID string) (Charge, error)
	FindByIDs(ctx context.Context, ids []string) ([]Charge, error)
	Delete(ctx context.Context, id string) error
}
