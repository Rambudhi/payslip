package model

import (
	"time"
)

type Overtime struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	Date      time.Time `gorm:"not null;uniqueIndex:idx_overtime_user_date"`
	Hours     float64   `gorm:"check:hours > 0 AND hours <= 3"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
