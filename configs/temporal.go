package configs

import (
	"github.com/pnnguyen58/aspire-code-challenge/pkg/clients"
	"time"
)

func LoadTempoConfig() *clients.TempoConfig {
	return mockTempoConfig()
}

func mockTempoConfig() (tc *clients.TempoConfig) {
	tc = &clients.TempoConfig{
		HostPort: C.Server.TempoHost,
		Namespace: &clients.Namespace{
			Name:                             "aspire-code-challenge",
			WorkflowExecutionRetentionPeriod: 1720 * time.Hour,
		},
		Workflows: map[string]*clients.Workflow{},
	}
	tc.Workflows["loan-workflow"] = &clients.Workflow{
		TaskQueueName:      "loan-workflow",
		ExecutionTimeout:   60 * time.Second,
		RunTimeout:         60 * time.Second,
		TaskTimeout:        60 * time.Second,
		MaximumInterval:    60 * time.Second,
		InitialInterval:    60 * time.Second,
		BackoffCoefficient: 2.0,
	}
	tc.Workflows["repayment-workflow"] = &clients.Workflow{
		TaskQueueName:      "repayment-workflow",
		ExecutionTimeout:   60 * time.Second,
		RunTimeout:         60 * time.Second,
		TaskTimeout:        60 * time.Second,
		MaximumInterval:    60 * time.Second,
		InitialInterval:    60 * time.Second,
		BackoffCoefficient: 2.0,
	}
	return
}
