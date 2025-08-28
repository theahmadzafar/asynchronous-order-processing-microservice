package container

import (
	"asynchronous-order-processing-microservice/internal/constants"
	"asynchronous-order-processing-microservice/internal/services"
	"asynchronous-order-processing-microservice/internal/transport/http/handlers"
	"context"
	"sync"

	"github.com/sarulabs/di"
)

func BuildHandlers(ctx context.Context, wg *sync.WaitGroup) []di.Def {
	return []di.Def{
		{
			Name: constants.METAHANDLER,
			Build: func(ctn di.Container) (interface{}, error) {
				return handlers.NewMetaHandler(), nil
			},
		}, {
			Name: constants.ORDERHANDLER,
			Build: func(ctn di.Container) (interface{}, error) {
				os := ctn.Get(constants.ORDERSERVICE).(*services.OrderService)
				return handlers.NewOrderHandler(os), nil
			},
		},
	}
}
