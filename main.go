package main

import (
	"net/http"
	"rest_api_sederhana/dto"
	"rest_api_sederhana/internal/api"
	"rest_api_sederhana/internal/config"
	"rest_api_sederhana/internal/connection"
	"rest_api_sederhana/internal/repository"
	"rest_api_sederhana/internal/service"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// app.Use(): Mengaplikasikan middleware secara global ke semua rute.
// Middleware di Endpoint: Mengaplikasikan middleware hanya untuk endpoint tertentu, menggunakan method chaining.
// Grup Rute: Mengaplikasikan middleware ke grup rute, sehingga middleware hanya diterapkan pada rute dalam grup tersebut.
func main() {
	cnf := config.Get()

	db := connection.GetDatabase(&cnf.Database)

	app := fiber.New()

	// disii kita akan mencoba melindungi nya dengan token
	// cara membuat middleware
	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("api unauthorized"))
		},
	})

	customerRepository := repository.NewCustomer(db)
	userRepository := repository.NewUser(db)
	bookRepository := repository.NewBook(db)
	bookStockRepository := repository.NewBookStock(db)
	journalRepository := repository.NewJournal(db)
	mediaRepository := repository.NewMedia(db)
	chargeRepository := repository.NewCharge(db)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	bookService := service.NewBook(bookRepository, bookStockRepository, mediaRepository, cnf)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)
	journalService := service.NewJournal(journalRepository, bookStockRepository, bookRepository, customerRepository, chargeRepository)
	mediaService := service.NewMedia(cnf, mediaRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, jwtMidd)
	api.NewBookStock(app, bookStockService, jwtMidd)
	api.NewJournal(app, journalService, jwtMidd)
	api.NewMediaApi(app, mediaService, cnf, jwtMidd)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
