package container

import (
	"asynchronous-order-processing-microservice/internal/constants"
	"asynchronous-order-processing-microservice/internal/services"
	"asynchronous-order-processing-microservice/pkg/nats"
	"context"
	"sync"

	"github.com/sarulabs/di"
)

func BuildServices(ctx context.Context, wg *sync.WaitGroup) []di.Def {
	return []di.Def{
		{
			Name: constants.ORDERSERVICE,
			Build: func(ctn di.Container) (interface{}, error) {
				n := ctn.Get(constants.QUEUES).(*nats.Nats)

				return services.NewOrderService(n), nil
			},
		},
	}
}
