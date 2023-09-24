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
	repayment := &models.Repayment{}
	for _, r := range unpaidRepayments {
		if amount >= r.Amount {
			repayment = r
			break
		}
	}
	if repayment.ID == 0 {
		return nil, fmt.Errorf("unacceptable repayment amount")
	}

	err = repositories.W.RepaymentRepo.UpdateState(ctx, repayment.ID, defined.PAID)
	if err != nil {
		return nil, err
	}
	log := &models.RepaymentLog{
		LoanID: req.LoanId,
		Amount: amount,
	}
	_ = repositories.W.RepaymentRepo.CreateLog(ctx, log)

	// TODO check repayment amount exceed
	if amount > repayment.Amount {
		prepay := &models.RepaymentPrepay{
			LoanID: repayment.LoanID,
			Amount: amount - repayment.Amount,
		}
		_ = repositories.W.RepaymentRepo.CreatePrepay(ctx, prepay)
	}
	// Update loan to PAID when all repayment paid
	if len(unpaidRepayments) == 1 {
		err = repositories.W.LoanRepo.UpdateState(ctx, repayment.LoanID, defined.PAID)
		if err != nil {
			return nil, err
		}
	}
	return &protogen.RepaymentCreateResponse{
		Data: &protogen.Repayment{
			Id:        repayment.ID,
			LoanId:    repayment.LoanID,
			Amount:    float32(repayment.Amount),
			State:     string(defined.PAID),
			DueDate:   timestamppb.New(repayment.DueDate),
			CreatedAt: timestamppb.New(repayment.CreatedAt),
			UpdatedAt: timestamppb.New(repayment.UpdatedAt),
		},
	}, nil
}

func CreateRepaymentCompensation(ctx context.Context, req *protogen.RepaymentCreateRequest,
) (*protogen.RepaymentCreateResponse, error) {
	return nil, nil
}
