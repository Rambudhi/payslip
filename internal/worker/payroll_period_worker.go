package worker

import (
	"fmt"

	"github.com/Rambudhi/payslip/internal/model"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/repository"
)

func RegisterPayrollPeriodWorker(w *Worker, repo repository.PayrollPeriodRepository) {
	w.RegisterHandler(queue.CreatePayrollPeriodJobName, func(data interface{}) error {
		log, ok := data.(model.PayrollPeriod)
		if !ok {
			return fmt.Errorf("invalid data type for payroll period")
		}
		return repo.Create(&log)
	})
}
