package routes

//add to cart
import (
	"golangsidang/golangsidang/repository"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App, r *repository.CartRepository) {
	api := app.Group("/api")
	api.Post("/addtocart", r.AddCourseToCart)
	api.Get("/getcart", r.GetCart)
}
