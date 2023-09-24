package activities

import (
	"context"

	"github.com/pnnguyen58/aspire-code-challenge/internal/defined"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
)

func ApproveLoan(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	err := repositories.W.LoanRepo.UpdateState(ctx, req.Id, defined.APPROVED)
	if err != nil {
		return nil, err
	}
	err = repositories.W.RepaymentRepo.UpdateStateByLoanID(ctx, req.Id, defined.APPROVED)
	if err != nil {
		return nil, err
	}
	return &protogen.LoanApproveResponse{
		Data: req.Id,
	}, nil
}

func ApproveLoanCompensation(ctx context.Context, req *protogen.LoanApproveRequest) (*protogen.LoanApproveResponse, error) {
	// TODO: implement approve loan compensate
	return nil, nil
}
