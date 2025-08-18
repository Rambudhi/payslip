package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Rambudhi/payslip/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	hashPassword := func(password string) string {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(hashed)
	}

	admin := model.User{
		Username: "admin",
		Password: hashPassword("admin"),
		Role:     "admin",
		Salary:   0,
	}
	db.Where(model.User{Username: "admin"}).FirstOrCreate(&admin)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		emp := model.User{
			Username: fmt.Sprintf("employee%d", i),
			Password: hashPassword("password"),
			Role:     "employee",
			Salary:   float64(rand.Intn(5000000) + 5000000), // random 5jt - 10jt
		}
		db.Where(model.User{Username: emp.Username}).FirstOrCreate(&emp)
	}

	fmt.Println("✅ Seeder berhasil! 1 admin + 100 employees ditambahkan.")
}
