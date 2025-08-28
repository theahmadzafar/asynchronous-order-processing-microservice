package container

import (
	"asynchronous-order-processing-microservice/internal/config"
	"asynchronous-order-processing-microservice/internal/constants"
	"asynchronous-order-processing-microservice/internal/transport/http"
	"asynchronous-order-processing-microservice/pkg/database"
	"asynchronous-order-processing-microservice/pkg/nats"
	"asynchronous-order-processing-microservice/pkg/notifier"
	"context"
	"sync"

	"github.com/sarulabs/di"
)

var (
	container di.Container
	once      sync.Once
)

func Build(ctx context.Context, wg *sync.WaitGroup) di.Container {
	once.Do(func() {
		builder, _ := di.NewBuilder()
		defs := []di.Def{{
			Name: constants.CONFIG,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.New()
			},
		}, {
			Name: constants.SERVER,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(constants.CONFIG).(*config.Config)
				publicHandlers := []http.Handler{
					ctn.Get(constants.METAHANDLER).(http.Handler),
					ctn.Get(constants.ORDERHANDLER).(http.Handler),
				}
				return http.NewServer(ctx, wg, &cfg.Server, publicHandlers), nil
			},
		}, {
			Name: constants.QUEUES,
			Build: func(ctn di.Container) (interface{}, error) {
				return nats.New()
			},
		}, {
			Name: constants.DATABASE,
			Build: func(ctn di.Container) (interface{}, error) {
				return database.NewDatabaseImplementation(), nil
			},
		}, {
			Name: constants.NOTIFIER,
			Build: func(ctn di.Container) (interface{}, error) {
				return notifier.New(), nil
			},
		},
		}
		builder.Add(defs...)

		requirements := BuildRequirement(ctx, wg)
		services := BuildServices(ctx, wg)
		handlers := BuildHandlers(ctx, wg)

		builder.Add(handlers...)
		builder.Add(requirements...)
		builder.Add(services...)

		container = builder.Build()
	})

	return container
}
