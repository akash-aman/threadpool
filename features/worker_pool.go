package features

import (
	"fmt"
	"sync"
	"workerpool/contracts"
)

type FixedWorkerPool struct {
	TaskQueue   contracts.TaskQueue
	Concurrency int
	Wg          *sync.WaitGroup
}

func (fwp *FixedWorkerPool) NewWorkerPool() chan bool {
	done := make(chan bool)

	fwp.Wg.Add(fwp.Concurrency)
	for i := 0; i < fwp.Concurrency; i++ {
		go Worker(done, fwp.TaskQueue.Queue(), fwp.Wg, i)
	}

	return done
}

func Worker(done <-chan bool, taskQueue chan contracts.Task, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case task := <-taskQueue:
			fmt.Printf("Worker %d processing task for number %v\n", id, task)
			task.Process()
		}
	}
}
