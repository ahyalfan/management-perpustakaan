package dto

import "time"

type JournalData struct {
	ID         string       `json:"id"`
	BookStock  string       `json:"book_stock"`
	Book       BookData     `json:"book"`
	Customer   CustomerData `json:"customer"`
	Status     string       `json:"status"`
	BorrowedAt time.Time    `json:"borrowed_at"`
	ReturnedAt time.Time    `json:"returned_at"`
}

type CreateJournalRequest struct {
	BookStock  string `json:"book_stock"`
	CustomerId string `json:"customer_id" validate:"required"`
	BookId     string `json:"book_id" validate:"required"`
}

type ReturnJournalRequest struct {
	JournalId string `json:"journal_id"`
	UserId    string `json:"user_id" validate:"required"`
}
