package api

import (
	"context"
	"net/http"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, as domain.AuthService) {
	aa := authApi{
		authService: as,
	}

	app.Post("/auth", aa.Login)

}

func (aa *authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := aa.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}
