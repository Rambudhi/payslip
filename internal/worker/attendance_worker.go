package worker

import (
	"fmt"
	"time"

	"github.com/Rambudhi/payslip/internal/model"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/repository"
)

func RegisterAttendanceWorker(w *Worker, repo repository.AttendanceRepository) {
	w.RegisterHandler(queue.CreateAttendanceJobName, func(data interface{}) error {
		attendance, ok := data.(model.Attendance)
		if !ok {
			return fmt.Errorf("invalid data type for attendance")
		}

		exists, _ := repo.FindByUserAndDate(attendance.UserID, attendance.Date)
		if exists != nil {
			return fmt.Errorf("attendance already submitted for today")
		}

		attendance.CreatedAt = time.Now()
		return repo.Create(&attendance)
	})
}
