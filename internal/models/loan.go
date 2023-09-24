package models

import (
	"fmt"
	"gorm.io/gorm/clause"
	"time"
)

type Loan struct {
	ID            uint64  `json:"id,omitempty"  gorm:"<-:create;AUTO_INCREMENT;primaryKey;NOT NULL"`
	CustomerID    uint64  `json:"customerId,omitempty"  gorm:"<-:create;column:customer_id;type:bigint;NOT NULL"`
	RepaymentType string  `json:"repaymentType,omitempty"  gorm:"<-:create;column:repayment_type;NOT NULL"`
	Amount        float64 `json:"amount,omitempty" gorm:"<-:create;column:amount;NOT NULL"`
	Term          int32   `json:"term,omitempty" gorm:"<-:create;column:term;NOT NULL"`
	State         string  `json:"state,omitempty" gorm:"column:state;NOT NULL"`

	Customer Customer `json:"customer,omitempty" gorm:"foreignKey:customer_id"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"<-:create;column:created_at;NULL"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;NULL"`
}

func (Loan) TableName() string {
	return "loans"
}

func (Loan) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}

func (l Loan) ValidateCreate() error {
	if l.Amount <= 0 {
		return fmt.Errorf("missing loan amount")
	}
	if l.Term <= 0 {
		return fmt.Errorf("missing loan term")
	}
	return nil
}
