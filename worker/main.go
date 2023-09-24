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

//
//func main() {
//	ctx := context.Background()
//	//signalChan := make(chan os.Signal, 1)
//	//signal.Notify(signalChan)
//	//go func() {
//	//	// Wait for termination signal
//	//	<-signalChan
//	//	// Trigger cancellation of the context
//	//	cancel()
//	//	// Wait for goroutine to finish
//	//	fmt.Println("The workers terminated gracefully")
//	//}()
//	tempoHost := configs.GetEnv("TEMPO_HOST", "localhost:7233")
//	tempoNS := configs.GetEnv("TEMPO_NAMESPACE", "aspire-code-challenge")
//	tempoTasks := configs.GetEnv("TEMPO_TASK_QUEUE", "loan-workflow,repayment-workflow")
//	cl, err := client.Dial(client.Options{
//		HostPort:  tempoHost,
//		Namespace: tempoNS,
//	})
//	defer cl.Close()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	taskQueues := strings.Split(tempoTasks, ",")
//	loanRepo := repositories.NewLoan()
//	repaymentRepo := repositories.NewRepayment()
//	repositories.Wire(loanRepo, repaymentRepo)
//	register(cl, taskQueues)
//	// Wait for a signal to shut down the server
//	<-ctx.Done()
//}

func main() {
	configs.ReadConfig()
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(context.Background())
	//signalChan := make(chan os.Signal, 1)
	//signal.Notify(signalChan)
	//
	//go func() {
	//	// Wait for termination signal
	//	<-signalChan
	//	// Trigger cancellation of the context
	//	cancel()
	//	// Wait for goroutine to finish
	//	fmt.Println("The service terminated gracefully")
	//}()

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
			w.RegisterWorkflow(workflows.CreateLoanWorkflow)
			w.RegisterActivity(activities.CreateLoan)
			w.RegisterActivity(activities.CreateLoanCompensation)
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

//func register(cl client.Client, taskQueues []string) {
//	for _, taskQueueName := range taskQueues {
//		switch taskQueueName {
//		case "loan-workflow":
//			w := worker.New(cl, taskQueueName, worker.Options{})
//			w.RegisterWorkflow(workflows.CreateLoanWorkflow)
//			w.RegisterActivity(activities.CreateLoan)
//			w.RegisterActivity(activities.CreateLoanCompensation)
//			// TODO: add more workflows and activities
//			go func() {
//				err := w.Run(worker.InterruptCh())
//				if err != nil {
//					log.Fatalln("Unable to start worker", err)
//				}
//			}()
//		default:
//			// TODO: add more workers
//			log.Println("app not defined")
//		}
//	}
//}
