package models

import (
	"gorm.io/gorm/clause"
	"time"
)

type Customer struct {
	ID uint64 `json:"id,omitempty"  gorm:"<-:create;AUTO_INCREMENT;primaryKey;NOT NULL"`
	Name string `json:"name,omitempty"  gorm:"unique;column:name;NOT NULL"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"<-:create;column:created_at;NULL"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;NULL"`
}

func (Customer) TableName() string {
	return "customers"
}

func (Customer) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}, {Name: "name"}},
			DoNothing: true,
		},
	}
}
