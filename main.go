package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
)

func main() {
	// Create a connection to Database
	//Get a new logger from loggers module
	// logger := loggers.GetLogger()

	// Get a handler for Vehicles
	// vehicleHandler := handlers.NewVehicleHandler(logger)

	router := mux.NewRouter()

	//Add a server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		loggers.GetLogger().InfoLogger.Println("Starting a Server on port ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			loggers.GetLogger().ErrorLogger.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	loggers.GetLogger().WarnLogger.Println("Received Terminate Command", sig)
	loggers.GetLogger().WarnLogger.Println("Initiating Graceful Shutdown")

	tc, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	server.Shutdown(tc)
	loggers.GetLogger().WarnLogger.Println("Graceful Shutdown Complete")
}
