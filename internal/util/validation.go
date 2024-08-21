package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Vallidate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := make(map[string]string)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			// res[v.StructField()] = v.Error()
			// mleakukan transalte, sebeanrya ini cuma akal akalan sih
			res[v.StructField()] = translateTag(v)
		}
	}
	return res
}

// bisa membuat custom message
func translateTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required":
		return fmt.Sprintf("field %s wajib diisi", fd.Field())
	}
	return "validasi gagal"
}
