package model

import (
	"gorm.io/gorm"
)

type Items struct{
	Id string `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"nama_barang" form:"name" gorm:"not null"`
	Description string `json:"deskripsi" form:"description" gorm:"not null"`
	Stock int `json:"jumlah_stok" form:"stock" gorm:"not null"`
	Price int `json:"harga" form:"price" gorm:"not null"`
	Category string `json:"kategori" form:"category" gorm:"not null"`
}

type Admin struct{
	gorm.Model
	Username string `json:"username" form:"username" gorm:"not null"`
	Password string `json:"password" form:"password" gorm:"not null"`
}