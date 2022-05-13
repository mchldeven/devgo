package models

import "time"

type OrderPelanggan struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User         User
	UserID       string `gorm:"size:36;index"`
	Order        Order
	OrderID      string `gorm:"size:36;index"`
	NamaDepan    string `gorm:"size:100;not null"`
	NamaBelakang string `gorm:"size:100;not null"`
	KotaID       string `gorm:"size:100;"`
	ProvinsiID   string `gorm:"size:100;"`
	Alamat1     string `gorm:"size:100;"`
	Alamat2     string `gorm:"size:100;"`
	Phone        string `gorm:"size:50;"`
	Email        string `gorm:"size:100;"`
	KodePos     string `gorm:"size:100;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}