package notification

import (
	"asynchronous-order-processing-microservice/pkg/nats"
	"asynchronous-order-processing-microservice/pkg/notifier"
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	conf    Config
	nats    *nats.Nats
	wg      *sync.WaitGroup
	sigTerm chan struct{}
	n       notifier.Notifier
}

func NewNotificationService(conf Config, nats *nats.Nats, wg *sync.WaitGroup, n notifier.Notifier) (*Notification, error) {
	v := &Notification{
		conf:    conf,
		nats:    nats,
		wg:      wg,
		sigTerm: make(chan struct{}),
		n:       n,
	}

	for i := 0; i < conf.WorkerCount; i++ {
		wg.Add(1)
		go v.worker(i)
	}

	return v, nil
}
func (v *Notification) Close() {
	time.Sleep(5 * time.Second)
	close(v.sigTerm)
}

func (v *Notification) worker(id int) {
	defer v.wg.Done()

	for {
		select {
		case <-v.sigTerm:
			fmt.Printf("worker %d shutting down\n", id)
			return
		default:
			// fmt.Printf("Notification worker %d running\n", id)
			// time.Sleep(1 * time.Second)
			cOrder := v.nats.NotifyQueue.Dequeue()
			if cOrder == nil {
				continue
			}

			v.n.Email()
		}
	}
}
