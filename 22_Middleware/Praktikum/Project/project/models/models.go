package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `gorm:"foreignkey:UserId"`
}

type UserResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}

type Blog struct {
	gorm.Model
	UserId int    `json:"user_id" form:"userid"`
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
}
