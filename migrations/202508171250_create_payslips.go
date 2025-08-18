package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreatePayslips(db *gorm.DB) {
	err := db.AutoMigrate(&model.Payslip{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: Payslip ✅")
}
