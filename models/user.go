package models

import (
	"mini-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Nama         string     `gorm:"type:varchar(255)" json:"nama"`
	KataSandi    string     `gorm:"type:varchar(255)" json:"kata_sandi"`
	NoTelp       string     `gorm:"type:varchar(255);uniqueIndex" json:"no_telp"`
	TanggalLahir *time.Time `gorm:"type:date" json:"tanggal_lahir"`
	JenisKelamin string     `gorm:"type:varchar(255)" json:"jenis_kelamin"`
	Tentang      string     `gorm:"type:text" json:"tentang"`
	Pekerjaan    string     `gorm:"type:varchar(255)" json:"pekerjaan"`
	Email        string     `gorm:"type:varchar(255)" json:"email"`
	IdProvinsi   string     `gorm:"type:varchar(255)" json:"id_provinsi"`
	IdKota       string     `gorm:"type:varchar(255)" json:"id_kota"`
	IsAdmin      bool       `gorm:"type:boolean" json:"IsAdmin"`
	Tokos        Toko
	Alamats      []Alamat
	Trxs         Trx
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.KataSandi = helpers.HashPass(u.KataSandi)
	err = nil
	return
}
