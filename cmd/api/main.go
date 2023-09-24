package main

import (
	"context"
	"fmt"
	"github.com/pnnguyen58/aspire-code-challenge/api/server"
	"github.com/pnnguyen58/aspire-code-challenge/api/server/handlers"
	"github.com/pnnguyen58/aspire-code-challenge/configs"
	"github.com/pnnguyen58/aspire-code-challenge/internal/app"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/clients"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/logger"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/persistence"
	"go.uber.org/fx"
	"os"
	"os/signal"
)

func main() {
	configs.ReadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)

	go func() {
		// Wait for termination signal
		<-signalChan
		// Trigger cancellation of the context
		cancel()
		// Wait for goroutine to finish
		fmt.Println("The service terminated gracefully")
	}()

	ap := fx.New(
		fx.Provide(
			logger.New,
			configs.LoadTempoConfig,
			configs.LoadDatabaseConfig,
			clients.NewTemporal,
			persistence.NewPostgresDB,
			context.TODO,
			// TODO add all providers

			app.NewLoan,
			repositories.NewLoan,
			repositories.NewRepayment,
			handlers.NewLoan,
		),
		fx.Invoke(
			repositories.Wire,
			server.ListenAndServe,
		),
	)
	if err := ap.Start(ctx); err != nil {
		os.Exit(1)
	}
}
