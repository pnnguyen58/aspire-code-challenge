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
	return l.app.Create(ctx, req)
}

func (l *Loan) ApproveLoan(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	return l.app.Approve(ctx, req)
}

func (l *Loan) GetLoan(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	return l.app.Get(ctx, req)
}

func (l *Loan) AddRepayment(ctx context.Context, req *protogen.RepaymentCreateRequest) (*protogen.RepaymentCreateResponse, error) {
	return l.app.AddRepayment(ctx, req)
}
