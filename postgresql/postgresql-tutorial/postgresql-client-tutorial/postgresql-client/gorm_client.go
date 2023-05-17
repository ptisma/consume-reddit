package main

import (
	"fmt"
	"log"

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
	return &PostgresConfig{DBHost: "127.0.0.1", DBUserName: "postgres", DBUserPassword: "StrongPassword", DBName: "reddit_db", DBPort: "5435"}

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

func CreateDatabase(config *PostgresConfig) error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBPort, config.DBUserName, config.DBUserPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", config.DBName)

	err = db.Exec(createDatabaseCommand).Error
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()

	return err
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	config := GetPostgresConfig()

	fmt.Println("config:", config)

	// err := CreateDatabase(config)
	// if err != nil {
	// 	log.Fatal("Failed to create the Database:", err)
	// }

	db, err := ConnectDB(config)
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}
	fmt.Println("? Connected Successfully to the Database")

	db.AutoMigrate(&Product{})

	err = db.Create(&Product{Code: "D44", Price: 100}).Error
	if err != nil {
		log.Fatal("Failed to create the product", err)
	}
	var product Product
	err = db.Last(&product).Error
	if err != nil {
		log.Fatal("Failed to fetch the product", err)
	}
	fmt.Println("Fetched product:", product)

}
