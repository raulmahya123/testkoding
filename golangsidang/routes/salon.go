// routes/routes.go
package routes

import (
	"golangsidang/golangsidang/repository"
	"golangsidang/golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func SalonRoutes(app *fiber.App, r *repository.SalonRepository) {
	api := app.Group("/api")
	api.Use(utils.TokenMiddleware)
	// Apply middleware for token verification to the remaining routes
	api.Post("/salon", r.CreateSalon)
}
