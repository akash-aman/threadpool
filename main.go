package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	feat "workerpool/features"
)

func main() {
	taskQueue := feat.NewChanTaskQueue(1000)
	wg := sync.WaitGroup{}
	workerPool := &feat.FixedWorkerPool{
		TaskQueue:   taskQueue,
		Concurrency: runtime.NumCPU(), // 20: for my machine. You can change it according to your machine.
		Wg:          &wg,
	}

	done := workerPool.NewWorkerPool()

	http.HandleFunc("/", handleRequest(workerPool))

	wg.Add(1)
	go StartServer(done, &wg)

	wg.Wait()
}

func handleRequest(workerPool *feat.FixedWorkerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		number, err := getNumberFromRequest(r)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		task := &feat.PrimeCheckTask{Number: number}
		if enqueueTask(workerPool, task) {
			fmt.Fprintf(w, "Task added: %d & Queue length is %d\n", number, len(workerPool.TaskQueue.Queue()))
		} else {
			http.Error(w, fmt.Sprintf("503 Service Unavailable: Task queue is full (current length: %d)", len(workerPool.TaskQueue.Queue())), http.StatusServiceUnavailable)
		}
	}
}

func getNumberFromRequest(r *http.Request) (int, error) {
	param := r.URL.Query().Get("number")
	return strconv.Atoi(param)
}

func enqueueTask(workerPool *feat.FixedWorkerPool, task *feat.PrimeCheckTask) bool {
	select {
	case workerPool.TaskQueue.Queue() <- task:
		return true
	default:
		return false
	}
}

func StartServer(done chan bool, wg *sync.WaitGroup) {
	defer close(done)
	defer wg.Done()
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
