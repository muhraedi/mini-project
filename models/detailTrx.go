package models

type DetailTrx struct {
	GormModel
	TrxID       uint
	LogProdukID uint
	TokoID      uint
	Kuantitas   int64 `gorm:"type:int" json:"kuantitas"`
	HargaTotal  int64 `gorm:"type:int" json:"harga_total"`
}
