package workerpool

type Pool interface {
	Make(numWorkers int)
	Handle(job func())
	Wait()
}

type Worker struct {
	ID   int
	Jobs int
}

type WPool struct {
	workers []*Worker
	pool    chan *Worker
}
