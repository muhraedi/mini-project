package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	NamaProduk    string `gorm:"type:varchar(255)" json:"nama_produk"`
	Slug          string `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller string `gorm:"type:varchar(255)" json:"harga_reseller"`
	HargaKonsumen string `gorm:"type:varchar(255)" json:"harga_konsumen"`
	Stok          int64  `gorm:"type:int" json:"stok"`
	Deskripsi     string `gorm:"type:text" json:"deskripsi"`
	TokoID        uint
	CategoryID    uint
	FotoProduks   []FotoProduk
	LogProduks    LogProduk
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
