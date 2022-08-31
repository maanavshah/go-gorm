package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	// "github.com/maanavshah/go-dummy-service/internal/model/book"
)

var DB *gorm.DB

func InitDb() {
	dbURL := "postgres://root:root@127.0.0.1:5432/testv1_golang"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	// Only uncomment this field during dev to see what SQL it is generating
	// Place that SQL query in a new file inside flyway migrations
	// db.AutoMigrate(&model.Book{})
	// Example log line: CREATE TABLE "books" ("id" bigserial,"title" text,"author" text,"desc" text,PRIMARY KEY ("id"))
	DB = db
}
