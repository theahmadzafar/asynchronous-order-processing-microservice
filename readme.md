

# Project: Asynchronous Order Processing Microservice


## Objective
Build a Go microservice that receives OrderCreated events and processes them through a pipeline with the following steps:

* Validation
* Persistence
* Notification

## Validation

status: complete.

to implement this validation i am getting the worker count from the config.yaml which can be set at start as this config is packaged into the image

internal/services/validation

## Persistence

status: complete.

to implement this Persistence i am getting the worker count from the config.yaml which can be set at start as this config is packaged into the image and Persistence is using the database which is simulated one and the core implementation can be changed because the Persistence service is dependant on interface not the core implementation

internal/services/Persistence

pkg/database

## validation

status: complete.

to implement this validation i am getting the worker count from the config.yaml which can be set at start as this config is packaged into the image and validation is using the notifier which is simulated one and the core implementation can be changed because the validation service is dependant on interface not the core implementation

internal/services/notification

pkg/notifier


## architecture of the project

project is built with the cloud native approaches using the dependency injection patterns and considering the single responsibility approach and hexagonal structure

* cmd/main.go project main dir
* config.yaml config file of the application
* dockerfile docker file is multi stage as required by the requirement
* internal/config is responsible for the loading and the parsing of the config
* internal/constants has all the constants declaration for the project dependency injection container
* internal/entities are the struct of the project nouns
* internal/services are the core services that are intracting and playing with the data this the logical hub of the project
* internal/transport is the http server of the project
* pkg has the external dependencies of the project like database and queues
## how to run

this project is configured with default parameters you can simply down load the mod by

```bash

go mod download
go run cmd/main.go

```
or you can make a build and run in your platform by dockerfile

```bash

docker buildx build . --progress=plain -t sample/order-processor-server:1.0.0 --no-cache  --build-arg TAG=1.0.0 --platform linux/amd64

```
