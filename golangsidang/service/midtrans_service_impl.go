package service

import (
	"fmt"
	"golangsidang/golangsidang/helper"
	"golangsidang/golangsidang/models"

	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	gomail "gopkg.in/gomail.v2"
)

type MidtransServiceImpl struct {
	Validate *validator.Validate
}

func NewMidtransServiceImpl(validate *validator.Validate) *MidtransServiceImpl {
	return &MidtransServiceImpl{
		Validate: validate,
	}
}

func (service *MidtransServiceImpl) Create(c *fiber.Ctx, request models.MidtransRequest) models.MidtransResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Initialize Midtrans snap client
	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// Convert user ID to string
	userID := strconv.Itoa(request.UserId)

	// Create customer address from the request
	custAddress := &midtrans.CustomerAddress{
		FName:       request.Customer.FirstName,
		LName:       request.Customer.LastName,
		Phone:       request.Customer.Phone,
		Address:     request.Customer.Address,
		City:        request.Customer.City,
		Postcode:    request.Customer.Postcode,
		CountryCode: request.Customer.CountryCode,
	}

	// Create transaction request for Midtrans
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + userID + "-" + request.ItemID,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    request.Customer.FirstName,
			LName:    request.Customer.LastName,
			Email:    request.Customer.Email,
			Phone:    request.Customer.Phone,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	// Create transaction
	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		helper.PanicIfError(errSnap.GetRawError())
	}

	// Send email notification
	err = service.sendEmailNotification(request.Customer.Email, response.RedirectURL)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Prepare response
	midtransResponse := models.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}

	return midtransResponse
}

func (service *MidtransServiceImpl) sendEmailNotification(toEmail string, paymentURL string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_USERNAME"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Payment Confirmation")
	m.SetBody("text/html", fmt.Sprintf("Thank you for your payment. Please complete your transaction by visiting the following URL: <a href='%s'>%s</a>", paymentURL, paymentURL))

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("GMAIL_USERNAME"), os.Getenv("GMAIL_PASSWORD"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email to %s: %v", toEmail, err)
		return err
	}

	log.Printf("Payment confirmation email sent to %s", toEmail)
	return nil
}
