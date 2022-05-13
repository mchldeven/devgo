package app

import (
	"github.com/michaeldeven/devgo/app/models"
)
type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Alamat{}},
		{Model: models.Product{}},
		{Model: models.GambarProduk{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderPelanggan{}},
		{Model: models.Payment{}},
		{Model: models.Shipment{}},
	}
}