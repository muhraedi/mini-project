package models

type Alamat struct {
	GormModel
	UserID       uint
	JudulAlamat  string `gorm:"type:varchar(255)" json:"judul_alamat"`
	NamaPenerima string `gorm:"type:varchar(255)" json:"nama_penerima"`
	NoTelp       string `gorm:"type:varchar(255)" json:"no_telp"`
	DetailAlamat string `gorm:"type:varchar(255)" json:"detail_alamat"`
	Trxs         Trx    `gorm:"foreignKey:AlamatPengiriman"`
}
