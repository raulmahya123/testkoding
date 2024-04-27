package models

import "gorm.io/gorm"

type Salon struct {
	Author        string `json:"author"`
	Title         string `json:"title"`
	Publisher     string `json:"publisher"`
	Delete_at     bool   `json:"delete_at"`
	Create_at     string `json:"create_at"`
	Create_by     string `json:"create_by"`
	Delete_by     string `json:"delete_by"`
	Update_at     string `json:"update_at"`
	Update_by     string `json:"update_by"`
	Salon_request string `json:"salon_request"`
	Is_active     bool   `json:"is_active"`
	Is_publish    bool   `json:"is_publish"`
}

func MigrateSalon(db *gorm.DB) error {
	err := db.AutoMigrate(&Salon{})
	return err
}
