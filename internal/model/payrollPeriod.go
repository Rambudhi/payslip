package model

import (
	"time"
)

type PayrollPeriod struct {
	ID        uint      `gorm:"primaryKey"`
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date;not null;check:end_date >= start_date"`
	Status    string    `gorm:"type:varchar(20);not null;check:status IN ('pending','processed');default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
