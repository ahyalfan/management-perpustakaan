package dto

type BookStockData struct {
	BookId string `json:"book_id"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

type CreateBookStockRequest struct {
	BookId string   `json:"book_id" validate:"required"`
	Codes  []string `json:"code" validate:"required"`
}

type DeleteBookStockRequest struct {
	Codes []string `json:"code" validate:"required;min=1;unique=true"`
}
