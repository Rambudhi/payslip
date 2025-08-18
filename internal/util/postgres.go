package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB
var Gorm *gorm.DB

func InitPostgres() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping Postgres: %v", err)
	}

	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the gorm")
	}

	sqlDB, err := gorm.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get generic DB from GORM: %v", err))
	}

	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	Gorm = gorm

	return gorm
}
