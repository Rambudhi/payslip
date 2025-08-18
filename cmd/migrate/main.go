package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rambudhi/payslip/migrations"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Println("Database connected, running migrations...")

	migrations.CreateUsers(db)
	migrations.CreateAttendances(db)
	migrations.CreateOvertimes(db)
	migrations.CreateReimbursements(db)
	migrations.CreatePayrollPeriods(db)
	migrations.CreatePayslips(db)
	migrations.CreateLogs(db)

	log.Println("✅ All migrations executed successfully!")
}
