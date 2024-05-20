package routes

import (
	service "golangsidang/golangsidang/Service"
	"golangsidang/golangsidang/controller"
	"golangsidang/golangsidang/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupPayment(app *fiber.App) {
	// Initialize dependencies
	validate := validator.New()
	midtransService := service.NewMidtransServiceImpl(validate)
	midtransController := controller.NewMidtransControllerImpl(midtransService)

	// Define routes
	api := app.Group("/api")
	api.Use(middleware.ErrorHandle()) // Apply error handling middleware

	midtrans := api.Group("/midtrans")
	midtrans.Post("/create", midtransController.Create)
}
