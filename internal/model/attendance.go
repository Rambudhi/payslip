package model

import (
	"time"
)

type Attendance struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	Date      time.Time `gorm:"not null;uniqueIndex:idx_attendance_user_date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
