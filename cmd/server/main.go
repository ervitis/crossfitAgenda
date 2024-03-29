package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ervitis/crossfitAgenda/credentials"
	"github.com/ervitis/crossfitAgenda/service/handlers/http"
	"github.com/ervitis/crossfitAgenda/service/usecases"
	"github.com/ervitis/crossfitAgenda/source_data"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGKILL)

	resourceManager := source_data.NewResourceManager(
		source_data.WithSourceDataClient(source_data.NewTwitterClient()),
	)
	srv := http.NewServer(http.NewHandler(usecases.New(resourceManager, credentials.New())))

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Panic(srv.Start(":8080"))
	}()

	go func() {
		<-done
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		defer wg.Done()

		log.Panic(srv.Shutdown(ctx))
	}()

	wg.Wait()
}
