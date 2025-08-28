package database

import "asynchronous-order-processing-microservice/internal/entities"

type databaseImplementation struct {
}

func NewDatabaseImplementation() Database {
	db := &databaseImplementation{}
	return db
}

func (d *databaseImplementation) Save(order entities.Order) error {
	return nil
}
