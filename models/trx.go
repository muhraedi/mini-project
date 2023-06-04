package models

type Trx struct {
	GormModel
	UserID           uint
	AlamatPengiriman uint
	HargaTotal       int64  `gorm:"type:int" json:"harga_total"`
	KodeInvoice      string `gorm:"type:varchar(255)" json:"kode_invoice"`
	MethodBayar      string `gorm:"type:varchar(255)" json:"method_Bayar"`
	DetailTrxs       DetailTrx
}
