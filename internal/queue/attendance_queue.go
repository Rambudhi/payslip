package queue

import "github.com/Rambudhi/payslip/internal/model"

const CreateAttendanceJobName = "attendance:create"

// EnqueueCreateAttendance push data attendance ke queue
func (q *Queue) EnqueueCreateAttendance(a model.Attendance) {
	q.Enqueue(Job{
		Name: CreateAttendanceJobName,
		Data: a,
	})
}
