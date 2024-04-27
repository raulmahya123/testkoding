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

type ProgramRepository struct {
	DB *gorm.DB
}

func NewProgramRepository(db *gorm.DB) *ProgramRepository {
	return &ProgramRepository{DB: db}
}

func init() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *ProgramRepository) CreateProgram(c *fiber.Ctx) error {
	programrequest := new(models.Program)
	if err := c.BodyParser(programrequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	// 1-3 level
	if programrequest.Level < 1 || programrequest.Level > 3 {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid level value",
			})
		return nil
	}

	programs := models.Program{
		Title:            programrequest.Title,
		Slug:             programrequest.Slug,
		Description:      programrequest.Description,
		Image_destop:     programrequest.Image_destop,
		Image_mobile:     programrequest.Image_mobile,
		Level:            programrequest.Level,
		Is_certification: programrequest.Is_certification,
		Url_Logo:         programrequest.Url_Logo,
		Pic_name:         programrequest.Pic_name,
		Pic_phone:        programrequest.Pic_phone,
		Start_at:         programrequest.Start_at,
		End_at:           programrequest.End_at,
		Is_active:        programrequest.Is_active,
		Is_publish:       programrequest.Is_publish,
		Create_at:        time.Now().Format("2006-01-02 15:04:05"),
		Delete:           false,
	}

	err := r.DB.Create(&programs).Error
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
			"data":   programs,
		})
	return nil
}

func (r *ProgramRepository) GetAllProgram(c *fiber.Ctx) error {
	var programs []models.Program
	r.DB.Find(&programs)
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   programs,
		})
	return nil
}

func (r *ProgramRepository) GetProgramById(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program
	r.DB.Where("id = ?", id).First(&program)
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   program,
		})
	return nil
}

func (r *ProgramRepository) UpdatedProgramById(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program
	r.DB.Find(&program, id)
	programRequest := new(models.Program)

	if err := c.BodyParser(programRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	program.Title = programRequest.Title
	program.Slug = programRequest.Slug
	program.Description = programRequest.Description
	program.Image_destop = programRequest.Image_destop
	program.Image_mobile = programRequest.Image_mobile
	program.Level = programRequest.Level
	program.Is_certification = programRequest.Is_certification
	program.Url_Logo = programRequest.Url_Logo
	program.Pic_name = programRequest.Pic_name
	program.Pic_phone = programRequest.Pic_phone
	program.Start_at = programRequest.Start_at
	program.End_at = programRequest.End_at
	program.Is_active = programRequest.Is_active
	program.Is_publish = programRequest.Is_publish
	program.Update_at = time.Now().Format("2006-01-02 15:04:05")
	err := r.DB.Save(&program).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   program,
		})
	return nil
}

func (r *ProgramRepository) DeleteProgramByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var program models.Program
	r.DB.Find(&program, id)
	program.Delete = true
	program.Delete_at = time.Now().Format("2006-01-02 15:04:05")
	err := r.DB.Save(&program).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   program,
		})
	return nil
}
