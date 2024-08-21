package domain

import "errors"

var ErrBookNotFound = errors.New("Book not found")
var ErrCustomerNotFound = errors.New("Customer not found")
var ErrBookStockNotFound = errors.New("Book stock not found")
var ErrBookStockAlreadyBorrowed = errors.New("Book stock already borrowed")
var ErrBookStockNotBorrowed = errors.New("Book stock not borrowed")
var ErrBookStockNotAvailable = errors.New("Book stock not available")
var ErrBookStockNotEnough = errors.New("Book stock not enough")
var ErrJournalNotFound = errors.New("Journal not found")
