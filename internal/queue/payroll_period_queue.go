package queue

import "github.com/Rambudhi/payslip/internal/model"

const CreatePayrollPeriodJobName = "payrollPeriod:create_period"

func (q *Queue) EnqueueCreatePayrollPeriod(p model.PayrollPeriod) {
	q.Enqueue(Job{
		Name: CreatePayrollPeriodJobName,
		Data: p,
	})
}
