package repositories

import (
	"context"
	"github.com/pnnguyen58/aspire-code-challenge/internal/defined"
	"gorm.io/gorm"

	"github.com/pnnguyen58/aspire-code-challenge/internal/models"
)

type Loan interface {
	GetByID(context.Context, uint64) (*models.Loan, error)
	Create(context.Context, *models.Loan) error
	UpdateState(context.Context, uint64, defined.State) error
	Delete(context.Context, uint64) error
}

type loan struct {
	db *gorm.DB
}

func NewLoan(db *gorm.DB) Loan {
	return &loan{
		db: db,
	}
}

func (l *loan) GetByID(ctx context.Context, loadID uint64) (*models.Loan, error) {
	var result models.Loan
	err := l.db.WithContext(ctx).
		Model(models.Loan{ID: loadID}).First(&result).Error
	return &result, err
}

func (l *loan) Create(ctx context.Context, loan *models.Loan) error {
	err := l.db.Clauses(loan.InsertClause()...).
		Model(models.Loan{}).
		WithContext(ctx).
		Create(loan).Error
	return err
}

func (l *loan) UpdateState(ctx context.Context, loadID uint64, state defined.State) error {
	err := l.db.Model(models.Loan{}).
		WithContext(ctx).
		Save(&models.Loan{ID: loadID, State: string(state)}).Error
	return err
}

func (l *loan) Delete(ctx context.Context, id uint64) error {
	err := l.db.WithContext(ctx).
		Delete(&models.RepaymentPrepay{}, id).Error
	return err
}
