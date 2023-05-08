package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
	ServerPort     string
	ClientOrigin   string
}

func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{DBHost: "127.0.0.1", DBUserName: "postgres", DBUserPassword: "postgres", DBName: "test", DBPort: "5435"}

}

func ConnectDB(config *PostgresConfig) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Category  string    `gorm:"not null"`
}

func main() {

	fmt.Println("Hello world")

	config := GetPostgresConfig()

	db, err := ConnectDB(config)
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}
	fmt.Println("? Connected Successfully to the Database")

	db.AutoMigrate(&Post{})

	db.Create(&Post{Title: "Test1", Content: "Hello world!", Category: "Hot"})

	var posts []Post
	db.Where("category = ? AND created_at >= ? AND created_at <= ?", "Hot", "2023-05-07", "2023-05-08").Find(&posts)

	fmt.Println(posts)

}
