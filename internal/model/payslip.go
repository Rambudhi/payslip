package model

import (
	"time"
)

type Payslip struct {
	ID                 uint    `gorm:"primaryKey"`
	UserID             uint    `gorm:"not null;index:idx_payslip_user_period,unique"`
	PayrollPeriodID    uint    `gorm:"not null;index:idx_payslip_user_period,unique"`
	BaseSalary         float64 `gorm:"default:0"`
	AttendanceSalary   float64 `gorm:"default:0"`
	OvertimeSalary     float64 `gorm:"default:0"`
	ReimbursementTotal float64 `gorm:"default:0"`
	TotalTakeHome      float64 `gorm:"default:0"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
