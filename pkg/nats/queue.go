package nats

import (
	"asynchronous-order-processing-microservice/internal/entities"
	"sync"
)

type OrderQueue struct {
	mu    sync.Mutex
	Queue []entities.Order
}

func (q *OrderQueue) Dequeue() *entities.Order {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.Queue) == 0 {
		q.Queue = nil
		return nil
	}
	o := q.Queue[0]
	q.Queue = q.Queue[1:]
	return &o
}
func (q *OrderQueue) Enqueue(o entities.Order) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Queue = append(q.Queue, o)
}
