package repository

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"golangsidang/golangsidang/models"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) AddCourseToCart(ctx *fiber.Ctx) error {
	// Membuat objek baru untuk menyimpan data permintaan
	cartRequest := new(models.Cartt)
	if err := ctx.BodyParser(cartRequest); err != nil {
		ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	// Memastikan bahwa Items yang dikirim adalah valid
	if cartRequest.Items == nil || len(cartRequest.Items) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "error",
				"message": "Items cannot be empty",
			})
		return errors.New("items cannot be empty")
	}

	// Ambil data keranjang lama dari database berdasarkan ID, jika ada
	var existingCart models.Cartt
	if err := r.DB.Where("id = ?", cartRequest.ID).Preload("Items").First(&existingCart).Error; err != nil {
		// Jika tidak ada data lama, gunakan data baru
		existingCart = *cartRequest
	} else {
		// Jika ada data lama, gabungkan dengan data baru
		existingCart.Nama = cartRequest.Nama
		existingCart.Author = cartRequest.Author
		existingCart.DeletedBy = false
		existingCart.CreatedAt = time.Now()
		existingCart.Items = append(existingCart.Items, cartRequest.Items...)
	}

	// Menghitung jumlah keseluruhan item dalam keranjang
	totalItems := len(existingCart.Items)

	// Mengatur total item ke dalam objek keranjang
	existingCart.Total = totalItems

	// Menyimpan data keranjang (lama atau baru) ke dalam database
	if err := r.DB.Save(&existingCart).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{
				"status": "error",
			})
		return err
	}

	ctx.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status": "success",
			"data":   existingCart,
		})
	return nil
}

func (r *CartRepository) GetCart(ctx *fiber.Ctx) error {
	var carts []models.Cartt
	r.DB.Find(&carts)

	//bagian item keranjang ambil dari tabel item keranjang
	for i := range carts {
		var items []models.CartItem
		r.DB.Model(&carts[i]).Association("Items").Find(&items)
		carts[i].Items = items
	}

	ctx.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"status": "success",
			"data":   carts,
		})
	return nil
}
