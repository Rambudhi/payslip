package queue

import "github.com/Rambudhi/payslip/internal/model"

const LogActivityJobName = "log:activity"

func (q *Queue) EnqueueLogActivity(d model.Log) {
	q.Enqueue(Job{
		Name: LogActivityJobName,
		Data: d,
	})
}
