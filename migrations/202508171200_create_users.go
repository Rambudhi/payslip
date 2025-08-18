package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: users ✅")
}
