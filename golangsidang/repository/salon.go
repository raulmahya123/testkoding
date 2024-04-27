package repository

import (
	"golangsidang/golangsidang/models"
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type SalonRepository struct {
	DB *gorm.DB
}

func NewSalonRepository(db *gorm.DB) *SalonRepository {
	return &SalonRepository{DB: db}
}

func init() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *SalonRepository) CreateSalon(c *fiber.Ctx) error {
	salonrequest := new(models.Salons)
	if err := c.BodyParser(salonrequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	// salon request ada makeup, hair, dan nail
	if salonrequest.Salon_request != "makeup" && salonrequest.Salon_request != "hair" && salonrequest.Salon_request != "nail" {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid salon request",
			})
	}

	salons := models.Salons{
		Publisher:     salonrequest.Publisher,
		Author:        salonrequest.Author,
		Is_active:     salonrequest.Is_active,
		Is_publish:    salonrequest.Is_publish,
		Title:         salonrequest.Title,
		Salon_request: salonrequest.Salon_request,
		Create_at:     time.Now().Format("2006-01-02 15:04:05"),
	}

	err := r.DB.Create(&salons).Error
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	c.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status": "success",
			"data":   salons,
		})
	return nil
}
