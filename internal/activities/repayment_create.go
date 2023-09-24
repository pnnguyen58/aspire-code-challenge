package activities

import (
	"context"
	"github.com/pnnguyen58/aspire-code-challenge/internal/defined"
	"github.com/pnnguyen58/aspire-code-challenge/internal/models"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateRepayment(ctx context.Context, repaymentRepo repositories.Repayment, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	amount := float64(req.Amount)
	repayment := &models.Repayment{
		LoanID:  req.LoanId,
		Amount:  amount,
		DueDate: req.DueDate.AsTime(),
		State:   string(defined.PENDING),
	}
	if err := repayment.ValidateCreate(); err != nil {
		return nil, err
	}
	err := repaymentRepo.Create(ctx, repayment)
	if err != nil {
		return nil, err
	}

	return &protogen.RepaymentCreateResponse{
		Data: &protogen.Repayment{
			Id:        repayment.ID,
			LoanId:    repayment.LoanID,
			Amount:    float32(repayment.Amount),
			State:     repayment.State,
			DueDate:   timestamppb.New(repayment.DueDate),
			CreatedAt: timestamppb.New(repayment.CreatedAt),
		},
	}, nil
}

func CreateRepaymentCompensation(ctx context.Context, repaymentRepo repositories.Repayment, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	return nil, nil
}
