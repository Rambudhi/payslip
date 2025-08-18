package worker

import (
	"fmt"

	"github.com/Rambudhi/payslip/internal/queue"
)

type Worker struct {
	Queue    *queue.Queue
	Handlers map[string]func(interface{}) error
}

func NewWorker(q *queue.Queue) *Worker {
	return &Worker{
		Queue:    q,
		Handlers: make(map[string]func(interface{}) error),
	}
}

func (w *Worker) RegisterHandler(jobName string, handler func(interface{}) error) {
	w.Handlers[jobName] = handler
}

// Start worker dengan N goroutine
func (w *Worker) Start(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			for job := range w.Queue.Jobs {
				if handler, ok := w.Handlers[job.Name]; ok {
					if err := handler(job.Data); err != nil {
						fmt.Printf("[Worker %d] error: %v\n", id, err)
					}
				} else {
					fmt.Printf("[Worker %d] no handler for job: %s\n", id, job.Name)
				}
			}
		}(i)
	}
}
