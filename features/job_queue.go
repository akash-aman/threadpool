package features

import (
	"workerpool/contracts"
)

type ChanTaskQueue struct {
	taskQueue chan contracts.Task
}

func (ctq *ChanTaskQueue) Enqueue(task contracts.Task) {
	ctq.taskQueue <- task
}

func (ctq *ChanTaskQueue) Dequeue() contracts.Task {
	return <-ctq.taskQueue
}

func (ctq *ChanTaskQueue) Queue() chan contracts.Task {
	return ctq.taskQueue
}


func NewChanTaskQueue(len int) *ChanTaskQueue {

	if len != 0 && len > 0 {
		return &ChanTaskQueue{
			taskQueue: make(chan contracts.Task, len),
		}
	}

	return &ChanTaskQueue{
		taskQueue: make(chan contracts.Task),
	}
}
