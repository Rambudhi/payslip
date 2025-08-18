package repository

import (
	"time"

	"github.com/Rambudhi/payslip/internal/model"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(a *model.Attendance) error
	FindByUserAndDate(userID uint, date time.Time) (*model.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) Create(a *model.Attendance) error {
	return r.db.Create(a).Error
}

func (r *attendanceRepository) FindByUserAndDate(userID uint, date time.Time) (*model.Attendance, error) {
	var attendance model.Attendance
	err := r.db.Where("user_id = ? AND DATE(date) = DATE(?)", userID, date).First(&attendance).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}
