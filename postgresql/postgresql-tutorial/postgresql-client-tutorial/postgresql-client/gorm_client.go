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
	return &PostgresConfig{DBHost: "127.0.0.1", DBUserName: "postgres", DBUserPassword: "password123", DBName: "golang-gorm", DBPort: "6500", ServerPort: "8000", ClientOrigin: "http://localhost:3000"}

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

func main() {

	fmt.Println("Hello world")

	config := GetPostgresConfig()

	db, err := ConnectDB(config)
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}
	fmt.Println("? Connected Successfully to the Database")

}
