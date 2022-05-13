package models

import "time"

type Alamat struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Nama       string `gorm:"size:100"`
	IsPrimary  bool
	KotaID     string `gorm:"size:100"`
	ProvinwoID string `gorm:"size:100"`
	Alamat1   string `gorm:"size:255"`
	Alamat2   string `gorm:"size:255"`
	Phone      string `gorm:"size:100"`
	Email      string `gorm:"size:100"`
	KodePos   string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}