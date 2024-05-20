package middleware

import (
	"errors"
	"log"
	"net/http"

	"golangsidang/golangsidang/helper"
	"golangsidang/golangsidang/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// RequestLogger middleware function for logging incoming requests
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Log request details
		log.Printf("Request: %s %s", c.Method(), c.Path())
		log.Printf("Request Body: %s", c.Body())

		// Continue processing the request
		return c.Next()
	}
}

// ErrorHandle middleware function for error handling
func ErrorHandle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				if validationErrors(c, r) {
					return
				}
				internalServerError(c, r)
			}
		}()
		return c.Next()
	}
}

// validationErrors function for handling validation errors
func validationErrors(c *fiber.Ctx, err interface{}) bool {
	if exception, ok := err.(validator.ValidationErrors); ok {
		var ve validator.ValidationErrors
		out := make([]models.ErrorResponse, len(exception))
		if errors.As(exception, &ve) {
			for _, fe := range ve {
				out = append(out, models.ErrorResponse{
					Field:   fe.Field(),
					Message: helper.MessageForTag(fe.Tag()),
				})
			}
		}
		webResponse := models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   out,
		}
		c.Status(http.StatusBadRequest).JSON(webResponse)
		return true
	}
	return false
}

// internalServerError function for handling internal server errors
func internalServerError(c *fiber.Ctx, err interface{}) {
	log.Printf("Internal Server Error: %v", err)
	webResponse := models.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	c.Status(http.StatusInternalServerError).JSON(webResponse)
}
