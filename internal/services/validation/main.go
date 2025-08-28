package validation

import (
	"asynchronous-order-processing-microservice/pkg/nats"
	"fmt"
	"sync"
	"time"
)

type Validation struct {
	conf    Config
	nats    *nats.Nats
	wg      *sync.WaitGroup
	sigTerm chan struct{}
}

func NewValidationService(conf Config, nats *nats.Nats, wg *sync.WaitGroup) (*Validation, error) {
	v := &Validation{
		conf:    conf,
		nats:    nats,
		wg:      wg,
		sigTerm: make(chan struct{}),
	}

	for i := 0; i < conf.WorkerCount; i++ {
		wg.Add(1)
		go v.worker(i)
	}

	return v, nil
}
func (v *Validation) Close() {
	time.Sleep(5 * time.Second)
	close(v.sigTerm)
}

func (v *Validation) worker(id int) {
	defer v.wg.Done()

	for {
		select {
		case <-v.sigTerm:
			fmt.Printf("worker %d shutting down\n", id)
			return
		default:
			// fmt.Printf("Validation worker %d running\n", id)
			// time.Sleep(1 * time.Second)
			cOrder := v.nats.ValidationQueue.Dequeue()
			if cOrder == nil {
				continue
			}

			v.nats.StorageQueue.Enqueue(*cOrder)
		}
	}
}
