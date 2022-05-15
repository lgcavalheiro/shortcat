package model

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
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

	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	user := viper.GetString("DB_USER")
	pwd := viper.GetString("DB_PWD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dbName)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("[SETUP] - DB connection error: %s", err.Error()))
	}

	err = db.AutoMigrate(&Shortcat{})
	if err != nil {
		log.Fatal(fmt.Sprintf("[MODEL MIGRATION] - Model Migration error: %s", err.Error()))
	}

	log.Println("[SETUP] - Migration complete.")
}
