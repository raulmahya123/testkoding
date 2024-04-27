package repository

import (
	"golangsidang/golangsidang/models"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type TugasRepository struct {
	DB *gorm.DB
}

func NewTugasRepository(db *gorm.DB) *SalonRepository {
	return &SalonRepository{DB: db}
}

func init() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *TugasRepository) CreateTugas(c *fiber.Ctx) error {
	tugasrequest := new(models.Tugas)
	if err := c.BodyParser(tugasrequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	tugas := models.Tugas{
		Tugas1:  tugasrequest.Tugas1,
		Tugas2:  tugasrequest.Tugas2,
		Tugas3:  tugasrequest.Tugas3,
		Tugas4:  tugasrequest.Tugas4,
		Tugas5:  tugasrequest.Tugas5,
		Tugas6:  tugasrequest.Tugas6,
		Tugas7:  tugasrequest.Tugas7,
		Tugas8:  tugasrequest.Tugas8,
		Tugas9:  tugasrequest.Tugas9,
		Tugas10: tugasrequest.Tugas10,
		Tugas11: tugasrequest.Tugas11,
		Tugas12: tugasrequest.Tugas12,
		Tugas13: tugasrequest.Tugas13,
		Tugas14: tugasrequest.Tugas14,
		Tugas15: tugasrequest.Tugas15,
		Tugas16: tugasrequest.Tugas16,
		Tugas17: tugasrequest.Tugas17,
		Status:  "false",
	}

	err := r.DB.Create(&tugas).Error
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
			"data":   tugas,
		})
	return nil
}
