package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreatePayrollPeriods(db *gorm.DB) {
	err := db.AutoMigrate(&model.PayrollPeriod{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: Payroll Period ✅")
}
