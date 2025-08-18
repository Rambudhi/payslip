package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreateOvertimes(db *gorm.DB) {
	err := db.AutoMigrate(&model.Overtime{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: Overtime ✅")
}
