package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/akash-scalent/gotodo/configs"
	"github.com/akash-scalent/gotodo/logger"
	"github.com/akash-scalent/gotodo/routes"
)

func main() {
	configs.InitializeConfiguration()
	logger := logger.InitializeLogger()

	s := http.Server{
		Addr:         fmt.Sprintf(":%d", configs.Config.SPort), // configure the bind address
		Handler:      *routes.NewRouter(logger).Handler,                        // set the default handler
		ReadTimeout:  5 * time.Second,                         // max time to read request from the client
		WriteTimeout: 10 * time.Second,                        // max time to write response to the client
		IdleTimeout:  120 * time.Second,                       // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Info().Msgf("Starting server on port %d", configs.Config.SPort)

		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal().Err(err).Msg("Error starting server")
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Info().Msgf("Got signal:%v", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	logger.Info().Msg("shutting down")
	os.Exit(0)
}
