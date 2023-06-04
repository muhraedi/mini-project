package models

type LogProduk struct {
	GormModel
	ProductID     uint
	NamaProduk    string `gorm:"type:varchar(255)" json:"nama_produk"`
	Slug          string `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller string `gorm:"type:varchar(255)" json:"harga_reseller"`
	HargaKonsumen string `gorm:"type:varchar(255)" json:"harga_konsumen"`
	Deskripsi     string `gorm:"type:text" json:"deskripsi"`
	TokoID        uint
	CategoryID    uint
	FotoProduks   []FotoProduk `gorm:"foreignKey:ProductID"`
	DetailTrxs    DetailTrx
}
