package repository

import (
	"golangsidang/golangsidang/models"
	"golangsidang/golangsidang/utils"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func CourseNewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func init() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (r *CourseRepository) CreateCourse(c *fiber.Ctx) error {
	courseRequest := new(models.Courses)
	if err := c.BodyParser(courseRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	// Check privacy value
	if courseRequest.Privacy != "PRIVACY" && courseRequest.Privacy != "PUBLIC" {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid privacy value",
			})
	}

	// ambil id dari user login

	// ambil id user dari token
	courses := models.Courses{
		Title:       courseRequest.Title,
		Slug:        courseRequest.Slug,
		Description: courseRequest.Description,
		Privacy:     courseRequest.Privacy,
		Image:       courseRequest.Image,
		Certificate: courseRequest.Certificate,
		Level:       courseRequest.Level,
		Price:       courseRequest.Price,
		Create_at:   time.Now().Format("2006-01-02 15:04:05"),
		Delete:      false,
		Status_enum: "ACTIVE",
	}

	err := r.DB.Create(&courses).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	c.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status": "success",
			"data":   courses,
		})
	return nil
}

func (r *CourseRepository) GetCourses(c *fiber.Ctx) error {
	var courses []models.Courses
	var count int64
	r.DB.Find(&courses)
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   courses,
		})

	r.DB.Model(&models.Courses{}).Count(&count)
	// Paginate the course
	r.DB.Find(&courses)
	pageNumber := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")

	pageNum, err := strconv.Atoi(pageNumber)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Invalid page number",
			})
	}
	pageSze, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid page size",
		})
	}
	// Paginate the users
	offset := (pageNum - 1) * pageSze
	r.DB.Offset(offset).Limit(pageSze).Find(&courses)

	totalPages, pageSizeNow := utils.Pagination(count, pageNum, pageSze)

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"course": courses,
			"meta": fiber.Map{
				"page": pageNum,

				"pageSize":   pageSze,
				"totalPages": totalPages,
				"totalItems": pageSizeNow,
			},
		},
	})
}

func (r *CourseRepository) GetCourseByid(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Courses
	r.DB.Find(&course, id)
	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   course,
		})
	return nil
}

func (r *CourseRepository) UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Courses
	r.DB.Find(&course, id)
	courseRequest := new(models.Courses)
	if err := c.BodyParser(courseRequest); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}
	course.Title = courseRequest.Title
	course.Slug = courseRequest.Slug
	course.Description = courseRequest.Description
	course.Privacy = courseRequest.Privacy
	course.Image = courseRequest.Image
	course.Certificate = courseRequest.Certificate
	course.Level = courseRequest.Level
	course.Price = courseRequest.Price
	course.Update_at = time.Now().Format("2006-01-02 15:04:05")
	course.Update_by = "admin"
	err := r.DB.Save(&course).Error
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
			"data":   course,
		})
	return nil
}

func (r *CourseRepository) DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Courses
	r.DB.Find(&course, id)
	course.Delete = true
	course.Delete_at = time.Now().Format("2006-01-02 15:04:05")
	course.Delete_by = "admin"
	err := r.DB.Save(&course).Error
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
			"data":   course,
		})
	return nil
}
