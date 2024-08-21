package api

import (
	"context"
	"path/filepath"
	"rest_api_sederhana/domain"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MediaApi struct {
	cnf          *config.Config
	mediaService domain.MediaService
}

func NewMediaApi(app *fiber.App, mediaService domain.MediaService, cnf *config.Config, authMid fiber.Handler) {
	ma := &MediaApi{mediaService: mediaService, cnf: cnf}

	v1 := app.Group("/api", authMid)
	v1.Post("/media", ma.Create)
	v1.Static("/media", cnf.Storage.BasePath)
}

// Create implements fiber.Handler
func (ma *MediaApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.cnf.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	media, err := ma.mediaService.Create(c, dto.CreatedMediaRequest{
		Path: filename,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(dto.CreateResponseSuccess(media))
}
