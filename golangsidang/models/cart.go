package models

import "time"

// CartItem adalah item dalam keranjang belanja
type CartItem struct {
	ID       uint    `json:"id"`
	CourseID uint    `json:"course_id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	CartID   uint    `json:"cart_id"` // Kunci asing untuk mengaitkan dengan Cartt
}

// Cart adalah model untuk keranjang belanja
type Cartt struct {
	ID        uint       `json:"id"`
	Nama      string     `json:"nama"`
	Author    string     `json:"author"`
	DeletedBy bool       `json:"deleted_by"`
	CreatedAt time.Time  `json:"created_at"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"` // Menetapkan kunci asing untuk relasi
	Total     int        `json:"total"`
}
