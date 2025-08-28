package main

import (
	"asynchronous-order-processing-microservice/internal/constants"
	"asynchronous-order-processing-microservice/internal/container"
	"asynchronous-order-processing-microservice/internal/services/notification"
	"asynchronous-order-processing-microservice/internal/services/persistance"
	"asynchronous-order-processing-microservice/internal/services/validation"
	"asynchronous-order-processing-microservice/internal/transport/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/net/context"
)

var wg sync.WaitGroup

func main() {
	ctx := context.Background()
	app := container.Build(ctx, &wg)

	server := app.Get(constants.SERVER).(*http.Server)

	go server.Run()

	validation := app.Get(constants.VALIDATION).(*validation.Validation)
	persistance := app.Get(constants.PERSISTENCE).(*persistance.Persistance)
	notification := app.Get(constants.NOTIFICATION).(*notification.Notification)

	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	<-termSig

	//shutting down system

	validation.Close()
	persistance.Close()
	notification.Close()
	server.Shutdown()

	wg.Wait()
}
