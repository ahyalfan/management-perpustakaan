package api

import (
	"context"
	"net/http"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type journalAPi struct {
	journalService domain.JournalService
}

func NewJournal(app *fiber.App, journalService domain.JournalService, authMid fiber.Handler) {
	ja := journalAPi{journalService: journalService}

	v1 := app.Group("/api", authMid)
	v1.Get("/journals", ja.index)
	v1.Post("/journals", ja.create)
	v1.Put("/journals/:id", ja.update)
	// v1.Delete("/journals/:id", ja.delete)
	// v1.Get("/journals/:id", ja.show)
}

func (ja *journalAPi) index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	customerId := ctx.Query("customer_id")
	status := ctx.Query("status")
	journals, err := ja.journalService.Index(c, domain.JournalSearch{
		CustomerId: customerId,
		Status:     status,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(journals))

}
func (ja *journalAPi) create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateJournalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	journal_id, err := ja.journalService.Create(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess("id : " + journal_id))

}

func (ja *journalAPi) update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	claim := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	err := ja.journalService.Return(c, dto.ReturnJournalRequest{
		JournalId: id,
		UserId:    claim["id"].(string),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess("success returned"))

}
