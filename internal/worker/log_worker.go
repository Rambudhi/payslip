package worker

import (
	"fmt"

	"github.com/Rambudhi/payslip/internal/model"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/repository"
)

func RegisterLogActivityWorker(w *Worker, repo repository.LogRepository) {
	w.RegisterHandler(queue.LogActivityJobName, func(data interface{}) error {
		period, ok := data.(model.Log)
		if !ok {
			return fmt.Errorf("invalid data type for log activity")
		}
		return repo.Create(&period)
	})
}
