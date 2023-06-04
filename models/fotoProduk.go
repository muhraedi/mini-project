package models

type FotoProduk struct {
	GormModel
	ProductID uint
	Url       string `gorm:"type:varchar(255)" json:"url"`
}
