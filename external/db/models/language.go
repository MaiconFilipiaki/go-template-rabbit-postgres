package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	ID    uint   `gorm:"primayKey, AUTO_INCREMENT, index"`
	Name  string `gorm:"column:name_language"`
	Books []Book `gorm:"many2many:book_languages;"`
}
