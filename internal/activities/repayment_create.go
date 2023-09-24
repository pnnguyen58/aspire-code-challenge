package activities

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pnnguyen58/aspire-code-challenge/internal/defined"
	"github.com/pnnguyen58/aspire-code-challenge/internal/models"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
)

func CreateRepayment(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	amount := roundAmount(float64(req.Amount))
	unpaidRepayments, err := repositories.W.RepaymentRepo.GetUnpaidByLoanID(ctx, req.LoanId)
	if err != nil {
		return nil, err
	}
	if len(unpaidRepayments) == 0 {
		return nil, fmt.Errorf("no unpaid repayment")
	}
	// TODO check repayment amount exceed
	if amount < unpaidRepayments[0].Amount {
		return nil, fmt.Errorf("unacceptable repayment amount")
	}

	err = repositories.W.RepaymentRepo.UpdateState(ctx, unpaidRepayments[0].ID, defined.PAID)
	if err != nil {
		return nil, err
	}
	log := &models.RepaymentLog{
		LoanID: req.LoanId,
		Amount: amount,
	}
	_ = repositories.W.RepaymentRepo.CreateLog(ctx, log)

	return &protogen.RepaymentCreateResponse{
		Data: &protogen.Repayment{
			Id:        unpaidRepayments[0].ID,
			LoanId:    unpaidRepayments[0].LoanID,
			Amount:    float32(unpaidRepayments[0].Amount),
			State:     string(defined.PAID),
			DueDate:   timestamppb.New(unpaidRepayments[0].DueDate),
			CreatedAt: timestamppb.New(unpaidRepayments[0].CreatedAt),
			UpdatedAt: timestamppb.New(unpaidRepayments[0].UpdatedAt),
		},
	}, nil
}

func CreateRepaymentCompensation(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	return nil, nil
}
