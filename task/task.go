package task

import "gorm.io/gorm"

func Run(db *gorm.DB) {
	SendEmailCron(db)
}
