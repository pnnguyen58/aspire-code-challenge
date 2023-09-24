package activities

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pnnguyen58/aspire-code-challenge/internal/repositories"
	protogen "github.com/pnnguyen58/aspire-code-challenge/pkg/proto_generated"
)

func GetLoan(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	loan, err := repositories.W.LoanRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	repayments, err := repositories.W.RepaymentRepo.GetByLoanID(ctx, req.Id)
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
			UpdatedAt: timestamppb.New(r.UpdatedAt),
		})
	}
	return &protogen.LoanGetResponse{
		Data: data,
	}, nil
}

func GetLoanCompensation(ctx context.Context, req *protogen.LoanGetRequest) (*protogen.LoanGetResponse, error) {
	// TODO: implement approve loan compensate
	return nil, nil
}
