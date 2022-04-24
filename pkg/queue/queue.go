package queue

import "sync"

type ConcurrentQueue struct {
	queue []string
	mu    sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	temp := q.queue[0]
	q.queue = q.queue[1:]
	return temp
}

func (q *ConcurrentQueue) Front() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.queue[0]
}

func (q *ConcurrentQueue) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue) == 0
}
