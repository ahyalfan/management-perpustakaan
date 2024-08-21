package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type ChargeRepository struct {
	db *gorm.DB
}

func NewCharge(conn *gorm.DB) domain.ChargeRepository {
	return &ChargeRepository{db: conn}
}

// Delete implements domain.ChargeRepository.
func (c *ChargeRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByIDs implements domain.ChargeRepository.
func (c *ChargeRepository) FindByIDs(ctx context.Context, ids []string) ([]domain.Charge, error) {
	panic("unimplemented")
}

// FindByJournalID implements domain.ChargeRepository.
func (c *ChargeRepository) FindByJournalID(ctx context.Context, journalID string) (domain.Charge, error) {
	panic("unimplemented")
}

// FindByUserID implements domain.ChargeRepository.
func (c *ChargeRepository) FindByUserID(ctx context.Context, userID string) ([]domain.Charge, error) {
	panic("unimplemented")
}

// Save implements domain.ChargeRepository.
func (c *ChargeRepository) Save(ctx context.Context, r *domain.Charge) (string, error) {
	err := c.db.WithContext(ctx).Create(r).Error
	if err != nil {
		return "", err
	}
	return r.ID, nil
}
