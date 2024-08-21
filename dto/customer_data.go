package dto

type CustomerData struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// kita bikin validation
type CreateCustomerRequest struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID string `json:"-"` // ini artinya dia akan diabaikan oleh jsonnya, tetapi kita bisa mengisinya manual
	// intinya sih user gak bisa otak atik gitu
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}
