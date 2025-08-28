package container

import (
	"asynchronous-order-processing-microservice/internal/config"
	"asynchronous-order-processing-microservice/internal/constants"
	"asynchronous-order-processing-microservice/internal/services/notification"
	"asynchronous-order-processing-microservice/internal/services/persistance"
	"asynchronous-order-processing-microservice/internal/services/validation"
	"asynchronous-order-processing-microservice/pkg/database"
	"asynchronous-order-processing-microservice/pkg/nats"
	"asynchronous-order-processing-microservice/pkg/notifier"
	"context"
	"sync"

	"github.com/sarulabs/di"
)

func BuildRequirement(ctx context.Context, wg *sync.WaitGroup) []di.Def {
	return []di.Def{
		{
			Name: constants.VALIDATION,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(constants.CONFIG).(*config.Config)
				n := ctn.Get(constants.QUEUES).(*nats.Nats)

				return validation.NewValidationService(cfg.Validation, n, wg)
			},
		}, {
			Name: constants.PERSISTENCE,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(constants.CONFIG).(*config.Config)
				n := ctn.Get(constants.QUEUES).(*nats.Nats)
				db := ctn.Get(constants.DATABASE).(database.Database)

				return persistance.NewPersistanceService(cfg.Persistance, n, wg, db)
			},
		}, {
			Name: constants.NOTIFICATION,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get(constants.CONFIG).(*config.Config)
				n := ctn.Get(constants.QUEUES).(*nats.Nats)
				notifier := ctn.Get(constants.NOTIFIER).(notifier.Notifier)

				return notification.NewNotificationService(cfg.Notification, n, wg, notifier)
			},
		},
	}
}
