package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type journalService struct {
	journalRepository   domain.JournalRepository
	bookStockRepository domain.BookStockRepository
	bookRepository      domain.BookRepository
	customerRepository  domain.CustomerRepository
	chargeRepository    domain.ChargeRepository
}

func NewJournal(journalRepository domain.JournalRepository,
	bookStockRepository domain.BookStockRepository,
	bookRepository domain.BookRepository,
	cutomerRepository domain.CustomerRepository,
	chargeRepository domain.ChargeRepository) domain.JournalService {
	return &journalService{
		journalRepository:   journalRepository,
		bookStockRepository: bookStockRepository,
		bookRepository:      bookRepository,
		customerRepository:  cutomerRepository,
		chargeRepository:    chargeRepository,
	}
}

// Create implements domain.journalService.
func (j *journalService) Create(ctx context.Context, req dto.CreateJournalRequest) (string, error) {
	_, err := j.bookStockRepository.FindByBookId(ctx, req.BookId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", domain.ErrBookNotFound
	}
	if err != nil {
		return "", err
	}
	stock, err := j.bookStockRepository.FindByCode(ctx, req.BookStock)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("book stock not found")
	}
	if stock.Status == domain.BookStockStatusBorrowed {
		return "", errors.New("book stock is already borrowed")
	}
	journal := domain.Journal{
		ID:         uuid.NewString(),
		CustomerId: req.CustomerId,
		BookID:     req.BookId,
		StockCode:  req.BookStock,
		Status:     domain.JournalStatusInProgress,
		DueAt:      sql.NullTime{Valid: true, Time: time.Now().AddDate(0, 0, 7)},
		BorrowedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	stock.Status = domain.BookStockStatusBorrowed
	stock.BorrowedAt = journal.BorrowedAt
	stock.BorrowerId = sql.NullString{Valid: true, String: journal.CustomerId}
	err = j.bookStockRepository.Update(ctx, &stock)
	if err != nil {
		return "", err
	}
	return j.journalRepository.Save(ctx, &journal)
}

// Index implements domain.journalService.
func (j *journalService) Index(ctx context.Context, req domain.JournalSearch) ([]dto.JournalData, error) {
	jornals, err := j.journalRepository.Find(ctx, req)
	log.Println(jornals)
	if err != nil {
		return nil, err
	}
	customerId := make([]string, 0)
	bookId := make([]string, 0)
	for _, ja := range jornals {
		customerId = append(customerId, ja.CustomerId)
		bookId = append(bookId, ja.BookID)
	}
	customers := make(map[string]domain.Customer)
	if len(customerId) > 0 {
		customerDB, err := j.customerRepository.FindByIds(ctx, customerId)
		if err != nil {
			return nil, err
		}
		for _, c := range customerDB {
			customers[c.ID] = c
		}
	}
	books := make(map[string]domain.Book)
	if len(bookId) > 0 {
		bookDB, err := j.bookRepository.FindByIds(ctx, bookId)
		if err != nil {
			return nil, err
		}
		for _, b := range bookDB {
			books[b.ID] = b
		}
	}
	result := make([]dto.JournalData, 0)
	for _, ja := range jornals {
		book := dto.BookData{}
		if v2, ok := books[ja.BookID]; ok {
			book = dto.BookData{
				Id:          v2.ID,
				Title:       v2.Title,
				Description: v2.Description,
				Isbn:        v2.Isbn,
			}
		}
		customer := dto.CustomerData{}
		if v1, ok := customers[ja.CustomerId]; ok {
			customer = dto.CustomerData{
				ID:   v1.ID,
				Code: v1.Code,
				Name: v1.Name,
			}
		}
		result = append(result, dto.JournalData{
			ID:         ja.ID,
			BookStock:  ja.StockCode,
			Book:       book,
			Customer:   customer,
			Status:     ja.Status,
			BorrowedAt: ja.BorrowedAt.Time,
			ReturnedAt: ja.ReturnedAt.Time,
		})
	}
	return result, nil

}

// Return implements domain.journalService.
func (j *journalService) Return(ctx context.Context, req dto.ReturnJournalRequest) error {
	journal, err := j.journalRepository.FindById(ctx, req.JournalId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrJournalNotFound
	}
	if err != nil {
		return err
	}
	stock, err := j.bookStockRepository.FindByCode(ctx, journal.StockCode)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		journal.Status = domain.JournalStatusCompleted
		journal.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}
		return j.journalRepository.Update(ctx, &journal)
		// return domain.ErrBookStockNotFound
	}
	if err != nil {
		return err
	}
	if stock.BorrowerId.String != journal.CustomerId {
		return errors.New("buku ini bukan milik customer ini")
	}
	stock.Status = domain.BookStockStatusAvailable
	stock.BorrowedAt = sql.NullTime{Valid: false}
	stock.BorrowerId = sql.NullString{Valid: false}
	err = j.bookStockRepository.Update(ctx, &stock)
	if err != nil {
		return err
	}
	journal.Status = domain.JournalStatusCompleted
	journal.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}
	err = j.journalRepository.Update(ctx, &journal)
	if err != nil {
		return err
	}
	hoursLate := time.Now().Sub(journal.DueAt.Time).Hours() //kita pasatikan telat berapa jam
	// jika sudah lebih dari 24 jam akan di didenda
	if hoursLate >= 24 {
		daysLate := int(hoursLate / 24)
		charge := domain.Charge{
			ID:           uuid.NewString(),
			JournalID:    journal.ID,
			DaysLate:     daysLate,
			DailyLateFee: 5000,
			UserID:       req.UserId,
			Total:        5000 * daysLate, // biaya per hari denda
		}
		_, err := j.chargeRepository.Save(ctx, &charge)
		if err != nil {
			return err
		}
	}
	return err
}
