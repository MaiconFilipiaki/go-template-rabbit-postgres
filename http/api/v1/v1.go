package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golangNetHttp/external/db/config"
	"golangNetHttp/external/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

func createUser(db *gorm.DB) (userNew *models.User) {

	languages := []models.Language{{
		Name: "Ingles",
	}, {
		Name: "Portgues",
	}}
	roleAdmin := &models.Role{}
	db.First(&roleAdmin, "Name=?", "ADMIN")
	db.Create(&languages)
	userNew = &models.User{
		Name:  "Maicon",
		Age:   12,
		Role:  *roleAdmin,
		Email: "mdfilipiaki@gmail.com",
		Books: []models.Book{
			{
				Title:     "Livro 1",
				Value:     120.00,
				Languages: languages,
			},
			{
				Title:     "Livro 2",
				Value:     130.00,
				Languages: []models.Language{languages[1]},
			},
		},
	}
	db.Create(userNew)
	return
}

func HandlersV1Http(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		db := config.GetDB()
		userNew := createUser(db)
		b, err := json.Marshal(userNew)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, string(b))
	})

	server.GET("/languages", func(ctx *gin.Context) {
		db := config.GetDB()
		var languages []models.Language
		db.Preload("Books").Find(&languages)
		ctx.JSON(200, languages)
	})

	server.GET("/books", func(ctx *gin.Context) {
		db := config.GetDB()
		var books []models.Book
		db.Preload("Languages").Find(&books)
		ctx.JSON(200, books)
	})
	server.GET("/users", func(ctx *gin.Context) {
		db := config.GetDB()
		var users []models.User
		db.Preload("Books").Preload("Books.Languages").Preload(clause.Associations).Find(&users)
		ctx.JSON(200, users)
	})

	server.GET("/drawsql", func(ctx *gin.Context) {
		db := config.GetDB()
		var users []models.User
		db.Raw("select * from users u, books b, book_languages bl , languages l where b.user_refer = u.id and b.id = bl.book_id and bl.language_id = l.id").Scan(&users)
		ctx.JSON(200, users)
	})
}
