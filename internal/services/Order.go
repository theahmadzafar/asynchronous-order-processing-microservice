package services

import (
	"asynchronous-order-processing-microservice/internal/entities"
	"asynchronous-order-processing-microservice/pkg/nats"
	"context"
)

type OrderService struct {
	nats *nats.Nats
}

func NewOrderService(nats *nats.Nats) *OrderService {
	return &OrderService{
		nats: nats,
	}
}

func (s *OrderService) Create(ctx context.Context, o entities.Order) error {
	s.nats.ValidationQueue.Enqueue(o)
	return nil
}
