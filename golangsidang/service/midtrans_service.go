package service

import (
	"golangsidang/golangsidang/models"

	"github.com/gofiber/fiber/v2"
)

type MidtransService interface {
	Create(c *fiber.Ctx, request models.MidtransRequest) models.MidtransResponse
}
