package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"golangNetHttp/external/db/models"
	smtp_infra "golangNetHttp/infra/smtp"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func SendEmailCron(db *gorm.DB) {
	fmt.Println("START CRON SEND EMAIL " + time.Now().String())
	c := cron.New()
	c.AddFunc("@every 1m10s", func() {
		users := []models.User{}
		db.Preload("Books").Preload("Books.Languages").Preload(clause.Associations).Where("email IS NOT NULL").Find(&users)
		for _, user := range users {
			smtp_infra.SendMail(user.Email, "GO EMAIL", "Opa, tudo bem ?")
		}
	})
	go c.Start()
}
