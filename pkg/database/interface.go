package database

import "asynchronous-order-processing-microservice/internal/entities"

type Database interface {
	Save(order entities.Order) error
}
