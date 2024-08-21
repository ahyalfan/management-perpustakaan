package service

import (
	"context"
	"errors"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"

	"gorm.io/gorm"
)

type bookStockService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &bookStockService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

// Create implements domain.BookStockService.
func (b *bookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	_, err := b.bookRepository.FindById(ctx, req.BookId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrBookNotFound
	}
	if err != nil {
		return err
	}

	stocks := make([]domain.BookStock, 0)
	for _, code := range req.Codes {
		stock := domain.BookStock{
			BookId: req.BookId,
			Code:   code,
			Status: domain.BookStockStatusAvailable,
		}
		stocks = append(stocks, stock)
	}
	return b.bookStockRepository.Save(ctx, stocks)

}

// Delete implements domain.BookStockService.
func (b *bookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	for _, stock := range req.Codes {
		stock, err := b.bookStockRepository.FindByCode(ctx, stock)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrBookNotFound
		}
		if err != nil {
			return err
		}
		if stock.Status == domain.BookStockStatusBorrowed {
			return errors.New("bukunya masih dipakai orang")
		}
	}
	return b.bookStockRepository.DeleteByCodes(ctx, req.Codes)
}
