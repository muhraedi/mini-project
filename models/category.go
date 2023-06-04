package models

type Category struct {
	GormModel
	NamaCategory string `gorm:"type:varchar(255)" json:"nama_category"`
	Products     []Product
	LogProduks   LogProduk
}
