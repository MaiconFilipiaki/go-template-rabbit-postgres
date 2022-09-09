package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint `gorm:"primayKey, AUTO_INCREMENT, index"`
	Name  string
	Age   int32
	Email string
	Role  Role   `gorm:"foreignKey:UserId"`    // Um usuario pode ter 1 role
	Books []Book `gorm:"foreignKey:UserRefer"` // Um usuario pode ter N books
}
