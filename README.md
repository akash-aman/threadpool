# Worker ThreadPool

This code is written to understand how we can implement a worker threadpool in Go. It demonstrates the basic concepts of task queues, worker pools, and concurrent task processing. 

**Note:** This implementation is for educational purposes only and should not be used in production directly. It lacks the robustness, error handling, and optimizations required for a production-grade system.

<video src="https://github.com/user-attachments/assets/19b895d1-2fa3-41b6-81a9-f828f1411772" width="352" height="720"></video>

## Knowledge Requirements
To understand and work with this codebase, you should have a basic understanding of the following concepts:

1. **Go Programming Language**: Familiarity with Go syntax, data structures, and concurrency primitives (goroutines, channels, for loops with select statement).
2. **HTTP Servers**: Basic knowledge of how to set up and handle HTTP requests in Go.
3. **Concurrency**: Understanding of concurrent programming concepts, including worker pools and task queues, wait groups & mutex.
4. **Prime Number Checking**: Basic understanding of prime number algorithms (optional, for understanding the task being processed).


## Implementation Details

### Components

1. **Task Queue**: A channel-based queue to hold tasks.
2. **Fixed Worker Pool**: A pool of worker goroutines that process tasks from the queue.
3. **HTTP Server**: An HTTP server that accepts requests and enqueues tasks.

### Code Explanation

1. **Main Function**:
    - Initializes the task queue with a capacity of 100,000.
    - Creates a worker pool with a concurrency level based on the number of CPU cores.
    - Starts the worker pool.
    - Sets up the HTTP server to handle requests.
    - Waits for all goroutines to complete.

2. **HTTP Handler**:
    - Validates the request method.
    - Extracts the "number" parameter from the request.
    - Creates a new task and enqueues it.
    - Returns a response indicating whether the task was successfully added or if the queue is full.

3. **Helper Functions**:
    - `getNumberFromRequest`: Extracts and converts the "number" parameter from the request.
    - `enqueueTask`: Attempts to enqueue a task and returns a boolean indicating success or failure.
    - `StartServer`: Starts the HTTP server and handles server errors.

### How to Test

You can test the implementation by sending HTTP GET requests to the server. The server expects a "number" parameter in the query string, which it uses to create a task.

#### Example Requests

1. **Start the Server**:
    - Run the Go program to start the server on port 8080.

2. **Send Requests**:
    - Use `curl` or any HTTP client to send requests to the server.

```sh
# Send a request with a number parameter
curl "http://localhost:8080/?number=123456789"

# Send multiple requests in parallel
for i in {1..10}; do curl "http://localhost:8080/?number=$i" & done
```

#### Example Response

- **Successful Task Addition**:
    ```sh
    Task added: 123456789 & Queue length is 1
    ```

- **Queue Full**:
    ```sh
    503 Service Unavailable: Task queue is full (current length: 100000)
    ```

### Note

This implementation is for educational purposes only and should not be used in production directly. It lacks the robustness, error handling, and optimizations required for a production-grade system.
