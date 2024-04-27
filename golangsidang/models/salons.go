package models

import "gorm.io/gorm"

type Salons struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	Author        string `gorm:"not null" json:"author"`
	Title         string `gorm:"not null" json:"title"`
	Publisher     string `gorm:"not null" json:"publisher"`
	Delete_at     bool   `gorm:"default:false" json:"delete_at"`
	Create_at     string `gorm:"not" json:"create_at"`
	Create_by     string `gorm:"not null" json:"create_by"`
	Delete_by     string `gorm:"-" json:"delete_by"`
	Update_at     string `gorm:"not" json:"update_at"`
	Update_by     bool   `gorm:"default:false" json:"update_by"`
	Salon_request string `gorm:"not null" json:"salon_request,omitempty"`
	Is_active     bool   `gorm:"default:true" json:"is_active"`
	Is_publish    bool   `gorm:"default:true" json:"is_publish"`
}

func MigrateSalons(db *gorm.DB) error {
	err := db.AutoMigrate(&Salons{})
	return err
}
