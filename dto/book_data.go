package dto

type BookData struct {
	Id          string `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CoverUrl    string `json:"cover_url"`
}

type BookShowData struct {
	BookData BookData
	Stocks   []BookStockData `json:"stocks"`
}

type CreateBookRequest struct {
	Isbn        string `json:"isbn" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CoverId     string `json:"cover_id"`
}

type UpdateBookRequest struct {
	Id          string `json:"-" validate:"required"`
	Isbn        string `json:"isbn" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CoverId     string `json:"cover_id"`
}
