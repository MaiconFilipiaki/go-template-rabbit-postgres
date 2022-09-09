package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID     uint `gorm:"primayKey, AUTO_INCREMENT, index"`
	UserId uint
	Name   string
}
