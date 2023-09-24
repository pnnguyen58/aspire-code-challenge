package activities

import (
	"context"
	"github.com/pnnguyen58/aspire-code-challenge/internal/defined"
	"github.com/pnnguyen58/aspire-code-challenge/internal/models"
	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"time"
)

func CreateLoan(ctx context.Context, req *protogen.LoanCreateRequest) (*protogen.LoanCreateResponse, error) {
	amount := float64(req.Amount)
	loan := &models.Loan{
		CustomerID:    req.CustomerId,
		RepaymentType: string(defined.WEEKLY),
		Amount:        amount,
		Term:          req.Term,
		State:         string(defined.PENDING),
	}
	if err := loan.ValidateCreate(); err != nil {
		return nil, err
	}
	err := repositories.W.LoanRepo.Create(ctx, loan)
	if err != nil {
		return nil, err
	}
	repayments := makeRepayment(*loan)
	err = repositories.W.RepaymentRepo.CreateBatch(ctx, repayments, defined.BACTH_SIZE)
	if err != nil {
		return nil, err
	}

	data := &protogen.Loan{
		Id:            loan.ID,
		CustomerId:    loan.CustomerID,
		RepaymentType: loan.RepaymentType,
		Amount:        float32(loan.Amount),
		Term:          loan.Term,
		State:         loan.State,
		CreatedAt:     timestamppb.New(loan.CreatedAt),
		UpdatedAt:     timestamppb.New(loan.UpdatedAt),
	}
	for _, r := range repayments {
		data.Repayments = append(data.Repayments, &protogen.Repayment{
			Id:        r.ID,
			LoanId:    r.LoanID,
			Amount:    float32(r.Amount),
			State:     r.State,
			DueDate:   timestamppb.New(r.DueDate),
			CreatedAt: timestamppb.New(r.CreatedAt),
		})
	}
	return &protogen.LoanCreateResponse{
		Data: data,
	}, nil
}

func CreateLoanCompensation(ctx context.Context, input *models.Loan) (*protogen.LoanCreateResponse, error) {
	// TODO: implement create loan compensate
	return nil, nil
}

func makeRepayment(loan models.Loan) []*models.Repayment {
	var (
		repayments []*models.Repayment
		i          int32
	)
	amount := loan.Amount
	term := loan.Term
	ratio := math.Pow(10, float64(defined.PRECISION))
	termAmount := math.Round((amount/float64(term))*ratio) / ratio
	for i = 1; i <= term; i++ {
		r := &models.Repayment{
			LoanID:  loan.ID,
			DueDate: setRepaymentDueDate(loan.CreatedAt, i),
			State:   string(defined.PENDING),
		}
		if i == term {
			r.Amount = amount - termAmount*float64(term-1)
		} else {
			r.Amount = termAmount
		}
		repayments = append(repayments, r)
	}
	return repayments
}

func setRepaymentDueDate(createdAt time.Time, next int32) time.Time {
	return createdAt.AddDate(0, 0, int(next*7))
	// TODO: implement other types of repayment: monthly, daily, ...
}
