package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/internal/controller"
	"test/internal/store/cachestore"
	"test/internal/subscriber"
	"time"
)

var (
	stanClusterID = "test-cluster"
	clientID      = "client1"
	subject       = "subject1"
	durable       = "my-durable"

	driverName     = "postgres"
	dataSourceName = "user=postgres dbname=store password=123 sslmode=disable"

	addr = "localhost:8080"
)

func Run() {
	r := chi.NewRouter()

	cStore, err := cachestore.NewCacheStore(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	sc, _, err := subscriber.ConnectAndSubscribe(stanClusterID, clientID, subject, durable, cStore)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	controller.Build(r, cStore)

	server := &http.Server{Addr: addr, Handler: r}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		fmt.Printf("\nReceived an interrupt, closing connection...\n\n")
		sc.Close()
		//sb.Unsubscribe()

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
