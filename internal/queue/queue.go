package queue

type Job struct {
	Name string
	Data interface{}
}

type Queue struct {
	Jobs chan Job
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		Jobs: make(chan Job, bufferSize),
	}
}

func (q *Queue) Enqueue(job Job) {
	q.Jobs <- job
}
