package workerpool

func New(numWorkers int) *WPool {
	wp := &WPool{
		workers: make([]*Worker, numWorkers),
		pool:    make(chan *Worker, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		wp.workers[i] = &Worker{ID: i}
	}

	return wp
}
