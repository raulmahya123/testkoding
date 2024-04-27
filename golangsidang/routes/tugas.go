// routes/routes.go
package routes

import (
	"golangsidang/golangsidang/repository"
	"golangsidang/golangsidang/utils"

	"github.com/gofiber/fiber/v2"
)

func TugasRoutes(app *fiber.App, r *repository.TugasRepository) {
	api := app.Group("/api")
	// Apply middleware for token verification to the remaining routes
	api.Use(utils.TokenMiddleware)

	api.Post("/tugas", r.CreateTugas)
	// user
}
