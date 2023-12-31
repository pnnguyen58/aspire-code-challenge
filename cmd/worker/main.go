package main

import (
	"context"
	"github.com/pnnguyen58/aspire-code-challenge/api/server/handlers"
	"github.com/pnnguyen58/aspire-code-challenge/configs"
	"github.com/pnnguyen58/aspire-code-challenge/internal/activities"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	"github.com/pnnguyen58/aspire-code-challenge/internal/workflows"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/clients"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/logger"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/persistence"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"log"
	"os"
)

func main() {
	configs.ReadConfig()
	ctx := context.Background()
	ap := fx.New(
		fx.Provide(
			logger.New,
			configs.LoadTempoConfig,
			configs.LoadDatabaseConfig,
			clients.NewTemporal,
			persistence.NewPostgresDB,
			context.TODO,
			// TODO add all providers

			repositories.NewLoan,
			repositories.NewRepayment,
			handlers.NewLoan,
		),
		fx.Invoke(
			repositories.Wire,
			register,
		),
	)
	if err := ap.Start(ctx); err != nil {
		os.Exit(1)
	}
	<-ctx.Done()
}

func register(cl client.Client, conf *clients.TempoConfig) {
	for _, taskQueue := range conf.Workflows {
		switch taskQueue.TaskQueueName {
		case "loan-workflow":
			w := worker.New(cl, taskQueue.TaskQueueName, worker.Options{})
			w.RegisterWorkflow(workflows.CreateLoan)
			w.RegisterWorkflow(workflows.ApproveLoan)
			w.RegisterWorkflow(workflows.GetLoan)
			w.RegisterWorkflow(workflows.CreateRepayment)

			w.RegisterActivity(activities.CreateLoan)
			w.RegisterActivity(activities.CreateLoanCompensation)
			w.RegisterActivity(activities.ApproveLoan)
			w.RegisterActivity(activities.ApproveLoanCompensation)
			w.RegisterActivity(activities.GetLoan)
			w.RegisterActivity(activities.GetLoanCompensation)
			w.RegisterActivity(activities.CreateRepayment)
			w.RegisterActivity(activities.CreateRepaymentCompensation)
			// TODO: add more workflows and activities
			go func() {
				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("Unable to start worker", err)
				}
			}()
		default:
			// TODO: add more workers
			log.Println("app not defined")
		}
	}
}
