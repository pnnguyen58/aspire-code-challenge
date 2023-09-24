package app

import (
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	"github.com/pnnguyen58/aspire-code-challenge/internal/workflows"
	"github.com/pnnguyen58/aspire-code-challenge/pkg/clients"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
)

type Loan interface {
	CreateOrchestration(context.Context, *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error)
}

type loan struct {
	logger         *zap.Logger
	temporalClient client.Client
	tempoWorkflow  *clients.Workflow
	loanRepo       repositories.Loan
	repaymentRepo  repositories.Repayment
}

func NewLoan(logger *zap.Logger, cl client.Client, tcf *clients.TempoConfig,
	loanRepo repositories.Loan, repaymentRepo repositories.Repayment) Loan {
	return &loan{
		logger:         logger,
		temporalClient: cl,
		tempoWorkflow:  tcf.Workflows["loan-workflow"],
		loanRepo:       loanRepo,
		repaymentRepo:  repaymentRepo,
	}
}

// CreateOrchestration create new loan use case
func (l *loan) CreateOrchestration(ctx context.Context, req *protogen.LoanCreateRequest,
) (*protogen.LoanCreateResponse, error) {
	res, err := ExecuteWorkflow[*protogen.LoanCreateRequest, *protogen.LoanCreateResponse](
		ctx, l.logger, l.temporalClient, l.tempoWorkflow, workflows.CreateLoanWorkflow, req)
	if err != nil {
		l.logger.Error("execute create loan workflow fail", zap.Error(err))
		return nil, errors.New("create loan failed")
	}
	if res == nil {
		return new(protogen.LoanCreateResponse), nil
	}
	return res, nil
}
