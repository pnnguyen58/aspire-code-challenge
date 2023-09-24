package handlers

import (
	"context"
	"go.uber.org/zap"

	"github.com/pnnguyen58/aspire-code-challenge/internal/app"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
)

func NewLoan(logger *zap.Logger, app app.Loan) *Loan {
	return &Loan{logger: logger, app: app}
}

type Loan struct {
	protogen.UnimplementedLoanServiceServer
	logger *zap.Logger
	app    app.Loan
}

func (l *Loan) CreateLoan(ctx context.Context, req *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error) {
	return l.app.CreateOrchestration(ctx, req)
}
