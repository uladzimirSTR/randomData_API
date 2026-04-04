package workerpool

import (
	"log"
	"time"
)

func (wp *WPool) Make(numWorkers int) {
	for idx, w := range wp.workers {
		wp.pool <- w
		log.Printf("Worker %d is ready to work\n", w.ID)
		wp.workers[idx].ID = idx
	}
}

func (wp *WPool) Handle(job func()) {
	worker := <-wp.pool

	go func() {
		defer func() {
			time.Sleep(45 * time.Second) // Simulate some processing time
			worker.Jobs++
			wp.pool <- worker
		}()
		job()
	}()
}

func (wp *WPool) Wait() {
	for i := 0; i < len(wp.workers); i++ {
		<-wp.pool
	}
}
