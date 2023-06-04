package models

type Toko struct {
	GormModel
	UserID     uint
	NamaToko   string `gorm:"type:varchar(255)" json:"nama_toko" form:"nama_toko"`
	UrlFoto    string `gorm:"type:varchar(255)" json:"url_foto" form:"photo"`
	Products   []Product
	LogProduks LogProduk
	DetailTrxs []DetailTrx
}
