package models

import (
	"gorm.io/gorm/clause"
	"time"
)

type RepaymentLog struct {
	ID     uint64  `json:"id,omitempty"  gorm:"<-:create;AUTO_INCREMENT;primaryKey;NOT NULL"`
	LoanID uint64  `json:"loanID,omitempty"  gorm:"column:loan_id;type:bigint;NOT NULL"`
	Amount float64 `json:"amount,omitempty" gorm:"column:amount;NOT NULL"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"<-:create;column:created_at;NULL"`

	Loan Loan `json:"loan,omitempty" gorm:"foreignKey:loan_id"`
}

func (RepaymentLog) TableName() string {
	return "repayment_logs"
}

func (RepaymentLog) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}
