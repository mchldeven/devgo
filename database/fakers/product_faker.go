package fakers

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/michaeldeven/devgo/app/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := faker.Name()
	return &models.Product{
		ID:               uuid.New().String(),
		UserID:           user.ID,
		Sku:              slug.Make(name),
		Name:             name,
		Slug:             slug.Make(name),
		Harga:            decimal.NewFromFloat(fakePrice()),
		Stok:            rand.Intn(100),
		Berat:           decimal.NewFromFloat(rand.Float64()),
		DeskripsiPendek: faker.Paragraph(),
		Deskripsi:      faker.Paragraph(),
		Status:           1,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

// precision | a helper function to set precision of price
func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}