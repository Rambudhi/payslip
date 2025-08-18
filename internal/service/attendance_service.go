package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Rambudhi/payslip/internal/model"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/request"
)

type AttendanceService interface {
	SubmitAttendance(req request.SubmitAttendanceRequest, ipAddr, requestID string) error
}

type attendanceService struct {
	queue *queue.Queue
}

func NewAttendanceService(q *queue.Queue) AttendanceService {
	return &attendanceService{queue: q}
}

func (s *attendanceService) SubmitAttendance(req request.SubmitAttendanceRequest, ipAddr, requestID string) error {

	today := time.Now()
	if today.Weekday() == time.Saturday || today.Weekday() == time.Sunday {
		return errors.New("attendance not allowed on weekend")
	}

	attendance := model.Attendance{
		UserID: req.UserID,
		Date:   today,
	}

	s.queue.EnqueueCreateAttendance(attendance)

	userID := req.UserID

	log := model.Log{
		TableName:   "attendances",
		RecordID:    attendance.ID,
		Action:      "create",
		OldData:     "{}",
		NewData:     fmt.Sprintf("{\"user_id\":%d,\"date\":\"%s\"}", req.UserID, today.Format("2006-01-02")),
		PerformedBy: &userID,
		IPAddress:   ipAddr,
		RequestID:   requestID,
	}

	s.queue.EnqueueLogActivity(log)

	return nil
}
