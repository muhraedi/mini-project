package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/mini_project?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Category{}, &User{}, &Alamat{}, &Toko{}, &Product{}, &FotoProduk{}, &LogProduk{}, &Trx{}, &DetailTrx{},
	)
	DB = db
}
