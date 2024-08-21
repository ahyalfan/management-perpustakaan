package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type JournalRepository struct {
	db *gorm.DB
}

func NewJournal(conn *gorm.DB) domain.JournalRepository {
	return &JournalRepository{db: conn}
}

// Find implements domain.JournalRepository.
func (j *JournalRepository) Find(ctx context.Context, se domain.JournalSearch) (result []domain.Journal, err error) {
	err = j.db.WithContext(ctx).Where("customer_id = ?", se.CustomerId).Find(&result).Error
	if se.Status != "" {
		j.db.Where("status =?", se.Status).Find(&result)
	}
	return
}

// FindById implements domain.JournalRepository.
func (j *JournalRepository) FindById(ctx context.Context, id string) (result domain.Journal, err error) {
	err = j.db.WithContext(ctx).Where("id =?", id).First(&result).Error
	return
}

// Save implements domain.JournalRepository.
func (j *JournalRepository) Save(ctx context.Context, journal *domain.Journal) (string, error) {
	err := j.db.WithContext(ctx).Create(&journal).Error
	if err != nil {
		return "", err
	}
	return journal.ID, nil
}

// Update implements domain.JournalRepository.
func (j *JournalRepository) Update(ctx context.Context, journal *domain.Journal) error {
	err := j.db.WithContext(ctx).Save(&journal).Error
	return err
}
