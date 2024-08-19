package contracts

type Task interface {
	Process()
}

type WorkerPool interface {
	Worker()
}

type TaskQueue interface {
	Enqueue(task Task)
	Dequeue() Task
	Queue() chan Task
}
