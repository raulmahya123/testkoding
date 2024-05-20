package models

type CustomerDetails struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required"`
	Address     string `json:"address" validate:"required"`
	City        string `json:"city" validate:"required"`
	Postcode    string `json:"postcode" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
}

type MidtransRequest struct {
	UserId   int             `json:"user_id" validate:"required"`
	ItemID   string          `json:"item_id" validate:"required"`
	ItemName string          `json:"item_name" validate:"required"`
	Amount   int64           `json:"amount" validate:"required"`
	Customer CustomerDetails `json:"customer" validate:"required"`
}
