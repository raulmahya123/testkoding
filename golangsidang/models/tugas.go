package models

import "gorm.io/gorm"

type Tugas struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	Tugas1        string `gorm:"not null" json:"tugas1"`
	Tugas2        string `gorm:"not null" json:"tugas2"`
	Tugas3        string `gorm:"not null" json:"tugas3"`
	Tugas4        string `gorm:"not null" json:"tugas4"`
	Tugas5        string `gorm:"not null" json:"tugas5"`
	Tugas6        string `gorm:"not null" json:"tugas6"`
	Tugas7        string `gorm:"not null" json:"tugas7"`
	Tugas8        string `gorm:"not null" json:"tugas8"`
	Tugas9        string `gorm:"not null" json:"tugas9"`
	Tugas10       string `gorm:"not null" json:"tugas10"`
	Tugas11       string `gorm:"not null" json:"tugas11"`
	Tugas12       string `gorm:"not null" json:"tugas12"`
	Tugas13       string `gorm:"not null" json:"tugas13"`
	Tugas14       string `gorm:"not null" json:"tugas14"`
	Tugas15       string `gorm:"not null" json:"tugas15"`
	Tugas16       string `gorm:"not null" json:"tugas16"`
	Tugas17       string `gorm:"not null" json:"tugas17"`
	Tugas18       string `gorm:"not null" json:"tugas18"`
	Salon_request string `gorm:"not null" json:"salon_request"`
	Status        string `gorm:"not null" json:"status"`
}

func MigrateTugas(db *gorm.DB) error {
	err := db.AutoMigrate(&Tugas{})
	return err
}
