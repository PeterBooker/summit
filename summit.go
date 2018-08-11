package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/PeterBooker/summit/internal/config"
	"github.com/PeterBooker/summit/internal/log"
	"github.com/PeterBooker/summit/internal/server"
)

func main() {
	// Create Logger
	l := log.New()

	// Create Config
	c := config.Setup(version, commit, date)

	// Setup server struct to hold all App data
	s := server.New(l, c)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Setup HTTP server.
	s.Setup()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
