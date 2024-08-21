package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type bookStockRepository struct {
	db *gorm.DB
}

func NewBookStock(conn *gorm.DB) domain.BookStockRepository {
	return &bookStockRepository{
		db: conn,
	}
}

// DeleteByBookId implements domain.BookStockRepository.
func (b *bookStockRepository) DeleteByBookId(ctx context.Context, id int) error {
	err := b.db.WithContext(ctx).Delete(&domain.BookStock{}, "book_id = ? ", id).Error
	return err
}

// DeleteByCodes implements domain.BookStockRepository.
func (b *bookStockRepository) DeleteByCodes(ctx context.Context, code []string) error {
	for _, c := range code {
		err := b.db.WithContext(ctx).Where("code =?", c).Delete(&domain.BookStock{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// FindByBookAndCode implements domain.BookStockRepository.
func (b *bookStockRepository) FindByCode(ctx context.Context, code string) (book domain.BookStock, err error) {
	err = b.db.WithContext(ctx).Where("code =?", code).Take(&book).Error
	return
}

// FindByBookId implements domain.BookStockRepository.
func (b *bookStockRepository) FindByBookId(ctx context.Context, id string) (books []domain.BookStock, err error) {
	err = b.db.WithContext(ctx).Where("book_id =?", id).Find(&books).Error
	return
}

// Save implements domain.BookStockRepository.
func (b *bookStockRepository) Save(ctx context.Context, bookStock []domain.BookStock) error {
	err := b.db.WithContext(ctx).Create(&bookStock).Error
	return err
}

// Update implements domain.BookStockRepository.
func (b *bookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	err := b.db.WithContext(ctx).Save(&stock).Error
	return err
}
