package models

import (
	"fmt"
	"gorm.io/gorm/clause"
	"time"
)

type Repayment struct {
	ID      uint64    `json:"id,omitempty"  gorm:"<-:create;AUTO_INCREMENT;primaryKey;NOT NULL"`
	LoanID  uint64    `json:"loanId,omitempty"  gorm:"<-:create;column:loan_id;type:bigint;NOT NULL"`
	Amount  float64   `json:"amount,omitempty" gorm:"<-:create;column:amount;NOT NULL"`
	DueDate time.Time `json:"dueDate,omitempty" gorm:"<-:create;column:due_date;NOT NULL"`
	State   string    `json:"state,omitempty" gorm:"column:state;NOT NULL"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"<-:create;column:created_at;NULL"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;NULL"`

	Loan Loan `json:"loan,omitempty" gorm:"foreignKey:loan_id;constraint:OnDelete:CASCADE;"`
}

func (Repayment) TableName() string {
	return "repayments"
}

func (Repayment) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}

func (r Repayment) ValidateCreate() error {
	if r.Amount <= 0 {
		return fmt.Errorf("missing repayment amount")
	}
	if r.DueDate.IsZero() {
		return fmt.Errorf("missing repayment due date")
	}
	if r.LoanID <= 0 {
		return fmt.Errorf("missing loan id")
	}
	return nil
}
