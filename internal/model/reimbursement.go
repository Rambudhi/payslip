package model

import (
	"time"
)

type Reimbursement struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null;index"`
	Amount      float64 `gorm:"not null;check:amount >= 0"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
