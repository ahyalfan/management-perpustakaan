package api

import (
	"context"
	"net/http"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, cs domain.CustomerService, authMid fiber.Handler) {
	ca := customerApi{
		customerService: cs,
	}

	// disini kita tambahkan handler di setiap methodnya, bisa pakai groouping
	v1 := app.Group("/api", authMid)

	// aatau jika mau ribet yg tinggal masukin satu persatu app.Get("/customers", authMid,ca.Index)
	v1.Get("/customers", ca.Index)
	v1.Post("/customers", ca.Create)
	v1.Put("/customers/:id", ca.Update)
	v1.Delete("/customers/:id", ca.Delete)
	v1.Get("/customers/:id", ca.Show)

}

func (ca *customerApi) Index(ctx *fiber.Ctx) error {
	// kita buat context
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (ca *customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	// kita bisa pakai parsing dari fiber
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	id, er := ca.customerService.Create(c, req)
	if er != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(er.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("created success, Id: " + id))

}

func (ca *customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	// kita bisa pakai parsing dari fiber
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	req.ID = ctx.Params("id")
	er := ca.customerService.Update(c, req)
	if er != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(er.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess("update success"))

}
func (ca *customerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	er := ca.customerService.Delete(c, id)
	if er != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(er.Error()))
	}

	return ctx.SendStatus(http.StatusOK)
}

func (ca *customerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	customer, er := ca.customerService.Show(c, id)
	if er != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(er.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(customer))
}
