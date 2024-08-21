package service

import (
	"context"
	"database/sql"
	"errors"
	"path"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookService struct {
	cnf                 *config.Config
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository     domain.MediaRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockR domain.BookStockRepository, mediaRepository domain.MediaRepository, snf *config.Config) domain.BookService {
	return &bookService{bookRepository: bookRepository, bookStockRepository: bookStockR, mediaRepository: mediaRepository, cnf: snf}
}

// Create implements domain.BookService.
func (b *bookService) Create(ctx context.Context, req dto.CreateBookRequest) (string, error) {
	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}
	book := domain.Book{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Isbn:        req.Isbn,
		CoverId:     coverId,
	}
	err := b.bookRepository.Save(ctx, &book)
	return book.ID, err
}

// Delete implements domain.BookService.
func (b *bookService) Delete(ctx context.Context, id string) error {
	_, err := b.bookRepository.FindById(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrBookNotFound
	}
	if err != nil {
		return err
	}
	// Cek stock dulu, jika ada yang masih ada, gagal delete
	bookStock, err := b.bookStockRepository.FindByBookId(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if len(bookStock) > 0 {
		return errors.New("cannot delete book, still has stock")
	}
	// Jika stock kosong, bisa delete
	return b.bookRepository.Delete(ctx, id)

}

// Index implements domain.BookService.
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrBookNotFound
	}
	if err != nil {
		return nil, err
	}
	coverId := make([]string, 0)
	for _, book := range books {
		if book.CoverId.Valid {
			coverId = append(coverId, book.CoverId.String)
		}
	}
	covers := make(map[string]string)
	if len(coverId) > 0 {
		coverDb, err := b.mediaRepository.FindByIds(ctx, coverId)
		if err != nil {
			return nil, err
		}
		for _, cover := range coverDb {
			covers[cover.ID] = path.Join(b.cnf.Server.Asset, cover.Path)
		}
	}
	var booksData []dto.BookData
	for _, book := range books {
		var coverUrl string
		if coverId, ok := covers[book.CoverId.String]; ok {
			coverUrl = coverId
		}
		booksData = append(booksData, dto.BookData{
			Id:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Isbn:        book.Isbn,
			CoverUrl:    coverUrl,
		})
	}
	return booksData, nil
}

// Show implements domain.BookService.
func (b *bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	book, err := b.bookRepository.FindById(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.BookShowData{}, domain.ErrBookNotFound
	}
	if err != nil {
		return dto.BookShowData{}, err
	}
	var coverUrl string
	if book.CoverId.Valid {
		cover, _ := b.mediaRepository.FindById(ctx, book.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(b.cnf.Server.Asset, cover.Path)
		}
	}
	bookData := dto.BookData{
		Id:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Isbn:        book.Isbn,
		CoverUrl:    coverUrl,
	}
	var bookStockData []dto.BookStockData
	stocks, er := b.bookStockRepository.FindByBookId(ctx, book.ID)

	if errors.Is(er, gorm.ErrRecordNotFound) {
		return dto.BookShowData{
			BookData: bookData,
			Stocks:   []dto.BookStockData{},
		}, domain.ErrBookNotFound
	}
	if er != nil {
		return dto.BookShowData{}, err
	}

	for _, stock := range stocks {
		bookStockData = append(bookStockData, dto.BookStockData{
			Code:   stock.Code,
			Status: stock.Status,
		})
	}

	return dto.BookShowData{
		BookData: bookData,
		Stocks:   bookStockData,
	}, nil
}

// Update implements domain.BookService.
func (b *bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	_, err := b.bookRepository.FindById(ctx, req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrBookNotFound
	}
	if err != nil {
		return err
	}
	// Cek stock dulu, jika ada yang masih ada, gagal delete
	bookStock, err := b.bookStockRepository.FindByBookId(ctx, req.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if len(bookStock) > 0 {
		return errors.New("cannot update book, still has stock")
	}

	coverId := sql.NullString{Valid: false, String: req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}
	// Jika stock kosong, bisa delete
	book := domain.Book{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Isbn:        req.Isbn,
		CoverId:     coverId,
	}
	return b.bookRepository.Update(ctx, &book)
}
