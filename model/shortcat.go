package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Shortcat struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Url       string    `json:"url" gorm:"not null"`
	ShortUrl  string    `json:"short_url" gorm:"not null;unique"`
	Clicks    uint64    `json:"clicks"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func Setup() {
	log.Println("[SETUP] - Starting migrations...")

	dsn := "root:example@tcp(127.0.0.1:3306)/shortcatdb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Shortcat{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[SETUP] - Migration complete.")
}
