package service

import (
	"fmt"
	"time"

	"github.com/Rambudhi/payslip/internal/model"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/request"
)

type PayrollPeriodService interface {
	Create(req request.CreatePayrollPeriodRequest, userID *uint, ipAddr, requestID string) error
}

type payrollPeriodService struct {
	queue *queue.Queue
}

func NewPayrollPeriodService(q *queue.Queue) PayrollPeriodService {
	return &payrollPeriodService{queue: q}
}

func (s *payrollPeriodService) Create(req request.CreatePayrollPeriodRequest, userID *uint, ipAddr, requestID string) error {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return err
	}

	period := model.PayrollPeriod{
		StartDate: startDate,
		EndDate:   endDate,
	}

	s.queue.EnqueueCreatePayrollPeriod(period)

	log := model.Log{
		TableName:   "payroll_periods",
		RecordID:    period.ID,
		Action:      "create",
		OldData:     "{}",
		NewData:     fmt.Sprintf("{\"start_date\":\"%s\",\"end_date\":\"%s\"}", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")),
		PerformedBy: userID,
		IPAddress:   ipAddr,
		RequestID:   requestID,
	}

	s.queue.EnqueueLogActivity(log)

	return nil
}
