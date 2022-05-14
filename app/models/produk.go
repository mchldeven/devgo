package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36;index"`
	User             User
	UserID           string `gorm:"size:36;index"`
	GambarProduks     []GambarProduk
	Categories       []Category      `gorm:"many2many:product_categories;"`
	Sku              string          `gorm:"size:100;index"`
	Nama             string          `gorm:"size:255"`
	Slug             string          `gorm:"size:255"`
	Harga            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stok             int
	Berat            decimal.Decimal `gorm:"type:decimal(10,2);"`
	DeskripsiPendek string          `gorm:"type:text"`
	Deskripsi      string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
