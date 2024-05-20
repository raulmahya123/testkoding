package controller

import (
	// "golangsidang/service"
	service "golangsidang/golangsidang/Service"
	"golangsidang/golangsidang/models"

	"github.com/gofiber/fiber/v2"
)

type MidtransControllerImpl struct {
	MidtransService service.MidtransService
}

func NewMidtransControllerImpl(midtransService service.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *fiber.Ctx) error {
	var request models.MidtransRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	// Call the Midtrans service to create a transaction
	midtransResponse := controller.MidtransService.Create(c, request)
	var webResponse models.WebResponse

	// Check if the response contains a token (indicating success)
	if midtransResponse.Token != "" {
		webResponse = models.WebResponse{
			Code:   fiber.StatusOK,
			Status: "Payment Success",
			Data:   midtransResponse,
		}
	} else {
		// If there's no token, consider it a failure
		webResponse = models.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Payment Failed",
			Data:   nil,
		}
	}

	return c.Status(webResponse.Code).JSON(webResponse)
}
