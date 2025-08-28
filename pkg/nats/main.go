package nats

type Nats struct {
	ValidationQueue OrderQueue
	StorageQueue    OrderQueue
	NotifyQueue     OrderQueue
}

func New() (*Nats, error) {
	n := &Nats{
		ValidationQueue: OrderQueue{},
		StorageQueue:    OrderQueue{},
		NotifyQueue:     OrderQueue{},
	}
	return n, nil
}
