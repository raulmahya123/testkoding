// main.go
package main

import (
	"log"
	"os"

	"golangsidang/golangsidang/models"

	"golangsidang/golangsidang/repository"
	"golangsidang/golangsidang/routes"
	"golangsidang/golangsidang/storage"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// added cors

	// Setup database connection
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("error connecting to the database")
	}

	// Migrate database models
	err = models.MigrateBooks(db)
	models.MigrateCourse(db)
	models.MigrateProgram(db)
	models.MigrateSalons(db)
	models.MigrateTugas(db)
	if err != nil {
		log.Fatal("error migrating books")
	}

	// Create a new Fiber app
	app := fiber.New()
	// CORS middleware
	app.Use(cors.New())
	repo2 := repository.Repositorry{db}
	course := repository.CourseRepository{db}
	repo3 := repository.ProgramRepository{db}
	salon := repository.SalonRepository{db}
	tugas := repository.TugasRepository{db}
	cart := repository.CartRepository{db}

	routes.UserRoutes(app, &repo2)
	routes.TugasRoutes(app, &tugas)
	routes.CourseRoutes(app, &course)
	routes.ProgramRoutes(app, &repo3)
	routes.SalonRoutes(app, &salon)
	routes.CartRoutes(app, &cart)
	routes.SetupPayment(app)

	// Create a repository instance using NewRepository
	repo := repository.NewRepository(db)

	routes.SetupRoutes(app, repo)

	// Start the server on port 3000 return sukes
	log.Fatal(app.Listen(":3000")) // Change the port if needed
	log.Println("Server is running on port 3000")

}
