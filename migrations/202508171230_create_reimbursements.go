package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreateReimbursements(db *gorm.DB) {
	err := db.AutoMigrate(&model.Reimbursement{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: Reimbursement ✅")
}
