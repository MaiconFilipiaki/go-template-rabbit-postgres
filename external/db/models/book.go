package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        uint   `gorm:"primayKey, AUTO_INCREMENT, index"`
	Title     string `gorm:"index"`
	Value     float64
	UserRefer uint
	Languages []Language `gorm:"many2many:book_languages;"` // Um Book Pode ter varias linguage e uma language pode ter varios livros N:N
}
