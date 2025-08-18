package migrations

import (
	"log"

	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

func CreateAttendances(db *gorm.DB) {
	err := db.AutoMigrate(&model.Attendance{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migrated: Attendance ✅")
}
