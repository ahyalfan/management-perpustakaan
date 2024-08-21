package repository

import (
	"context"
	"rest_api_sederhana/domain"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBook(conn *gorm.DB) domain.BookRepository {
	return &bookRepository{
		db: conn,
	}
}

// Delete implements domain.bookRepository.
func (b *bookRepository) Delete(ctx context.Context, id string) error {
	err := b.db.WithContext(ctx).Delete(&domain.Book{}, "id = ?", id).Error
	return err
}

// FindAll implements domain.bookRepository.
func (b *bookRepository) FindAll(ctx context.Context) (books []domain.Book, err error) {
	err = b.db.WithContext(ctx).Find(&books).Error
	return
}

// FindById implements domain.bookRepository.
func (b *bookRepository) FindById(ctx context.Context, id string) (book domain.Book, err error) {
	err = b.db.WithContext(ctx).Take(&book, "id = ?", id).Error
	return
}

// Save implements domain.bookRepository.
func (b *bookRepository) Save(ctx context.Context, book *domain.Book) error {
	err := b.db.WithContext(ctx).Create(book).Error
	return err
}

// Update implements domain.bookRepository.
func (b *bookRepository) Update(ctx context.Context, book *domain.Book) error {
	err := b.db.WithContext(ctx).Save(book).Error
	return err
}
func (b *bookRepository) FindByIds(ctx context.Context, id []string) (books []domain.Book, err error) {
	// pakai in  di sql
	err = b.db.WithContext(ctx).Where("id IN (?)", id).Find(&books).Error
	return

	// pakai perulangan manual
	// for i := 0; i < len(id); i++ {
	// 	var book domain.Book
	// 	err = b.db.WithContext(ctx).Where("id =?", id[i]).First(&book).Error
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	books = append(books, book)
	// }
	// return
}
