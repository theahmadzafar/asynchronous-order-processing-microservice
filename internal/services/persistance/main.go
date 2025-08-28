package persistance

import (
	"asynchronous-order-processing-microservice/pkg/database"
	"asynchronous-order-processing-microservice/pkg/nats"
	"fmt"
	"sync"
	"time"
)

type Persistance struct {
	conf    Config
	nats    *nats.Nats
	wg      *sync.WaitGroup
	sigTerm chan struct{}
	db      database.Database
}

func NewPersistanceService(conf Config, nats *nats.Nats, wg *sync.WaitGroup, db database.Database) (*Persistance, error) {
	v := &Persistance{
		conf:    conf,
		nats:    nats,
		wg:      wg,
		sigTerm: make(chan struct{}),
		db:      db,
	}

	for i := 0; i < conf.WorkerCount; i++ {
		wg.Add(1)
		go v.worker(i)
	}

	return v, nil
}
func (v *Persistance) Close() {
	time.Sleep(5 * time.Second)
	close(v.sigTerm)
}

func (v *Persistance) worker(id int) {
	defer v.wg.Done()

	for {
		select {
		case <-v.sigTerm:
			fmt.Printf("worker %d shutting down\n", id)
			return
		default:
			// fmt.Printf("Persistance worker %d running\n", id)
			// time.Sleep(1 * time.Second)
			cOrder := v.nats.StorageQueue.Dequeue()
			if cOrder == nil {
				continue
			}

			v.db.Save(*cOrder)

			v.nats.NotifyQueue.Enqueue(*cOrder)
		}
	}
}
